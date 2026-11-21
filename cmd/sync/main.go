package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/startracex/go-pub"
	"golang.org/x/mod/modfile"
)

const (
	upstreamRepo    = "https://github.com/microsoft/TypeScript.git"
	contentsURL     = "https://api.github.com/repos/microsoft/TypeScript/contents/tsc?ref=main"
	publishedModule = "github.com/startracex/tsgo"
	upstreamModule  = "github.com/microsoft/TypeScript/tsc"
	revFileName     = ".github/tsc-sync-rev"
	goPubModule     = "github.com/startracex/go-pub"
	goPubVersion    = "v0.0.0-20251012180452-cce541aafcd7"
)

// revKeys are the tsc inputs that affect published packages.
// testdata and cmd are intentionally absent: testdata is huge and is not
// copied, and cmd is the upstream CLI rather than the library.
var revKeys = []string{"internal", "go.mod", "go.sum", "LICENSE"}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "sync:", err)
		os.Exit(1)
	}
}

func run() error {
	root, err := findRoot()
	if err != nil {
		return err
	}
	remote, err := fetchRemoteRev()
	if err != nil {
		return err
	}
	local, err := readRevFile(filepath.Join(root, revFileName))
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if err == nil && remote.equal(local) {
		fmt.Printf("tsc unchanged (internal %s)\n", remote["internal"])
		return setChanged(false)
	}

	tmp, err := os.MkdirTemp("", "tsgo-upstream-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)

	fmt.Println("fetching tsc/internal (sparse, depth 1)")
	if err := cloneTSC(tmp); err != nil {
		return err
	}
	got, err := readGitRev(tmp)
	if err != nil {
		return err
	}
	tscDir := filepath.Join(tmp, "tsc")
	if err := publish(root, tscDir); err != nil {
		return err
	}
	if err := tidy(root); err != nil {
		return err
	}
	// Record the revision only after tidy succeeds, so a failed tidy is retried.
	if err := writeRevFile(filepath.Join(root, revFileName), got); err != nil {
		return err
	}
	fmt.Printf("synced tsc internal %s\n", got["internal"])
	return setChanged(true)
}

func findRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("go.mod not found from %s", dir)
		}
		dir = parent
	}
}

func publish(root, tscDir string) error {
	outDir := filepath.Join(root, "pkg")
	if err := os.RemoveAll(outDir); err != nil {
		return err
	}
	fmt.Println("publishing internal packages to pkg/")
	r, err := pub.NewModRoot(filepath.Join(tscDir, "go.mod"))
	if err != nil {
		return err
	}
	pc := &pub.PublicConfig{
		OutDir:         outDir,
		InternalOutDir: outDir,
		IncludeModFile: false,
		IncludeSumFile: false,
	}
	if err := pc.Public(r); err != nil {
		return err
	}
	if err := copyNonGo(filepath.Join(tscDir, "internal"), outDir); err != nil {
		return err
	}
	if err := rewriteImports(outDir); err != nil {
		return err
	}
	if _, err := os.Stat(filepath.Join(outDir, "internal")); err == nil {
		return fmt.Errorf("pkg/internal still exists; module-level internal was not lifted")
	}
	if err := copyFile(filepath.Join(tscDir, "LICENSE"), filepath.Join(root, "LICENSE")); err != nil {
		return err
	}
	if err := writeGoMod(root, filepath.Join(tscDir, "go.mod")); err != nil {
		return err
	}
	return copyFile(filepath.Join(tscDir, "go.sum"), filepath.Join(root, "go.sum"))
}

func copyNonGo(internalDir, outDir string) error {
	return filepath.WalkDir(internalDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || strings.HasSuffix(d.Name(), ".go") {
			return nil
		}
		rel, err := filepath.Rel(internalDir, path)
		if err != nil {
			return err
		}
		return copyFile(path, filepath.Join(outDir, rel))
	})
}

func rewriteImports(pkgDir string) error {
	return filepath.WalkDir(pkgDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		next, err := rewriteImportBytes(b)
		if err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}
		if bytes.Equal(b, next) {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		return os.WriteFile(path, next, info.Mode().Perm())
	})
}

func rewriteImportBytes(b []byte) ([]byte, error) {
	// Generators embed the upstream internal import as a quoted string.
	// Rewrite that before the already-lifted import paths so "/internal/"
	// is not left behind. Nested packages such as vfs/internal stay put.
	internalImport := []byte(`"` + upstreamModule + `/internal/`)
	publicPrefix := []byte(`"` + publishedModule + `/pkg/`)
	if bytes.Contains(b, internalImport) {
		b = bytes.ReplaceAll(b, internalImport, publicPrefix)
	}
	liftedImport := []byte(`"` + upstreamModule + `/`)
	if bytes.Contains(b, liftedImport) {
		b = bytes.ReplaceAll(b, liftedImport, publicPrefix)
	}
	if bytes.Contains(b, []byte(`"`+upstreamModule)) {
		return nil, fmt.Errorf("still imports %s", upstreamModule)
	}
	if bytes.Contains(b, []byte(`"`+publishedModule+`/pkg/internal/`)) {
		return nil, fmt.Errorf("import still points at module-level internal")
	}
	return b, nil
}

func writeGoMod(root, upstreamGoMod string) error {
	data, err := os.ReadFile(upstreamGoMod)
	if err != nil {
		return err
	}
	f, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return err
	}
	if f.Module == nil {
		return fmt.Errorf("upstream go.mod has no module statement")
	}
	// AddModuleStmt updates the syntax line Format prints. Assigning
	// Module.Mod.Path alone leaves the original module path in the file.
	if err := f.AddModuleStmt(publishedModule); err != nil {
		return err
	}
	if err := f.AddRequire(goPubModule, goPubVersion); err != nil {
		return err
	}
	f.Cleanup()
	out, err := f.Format()
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(root, "go.mod"), out, 0o644)
}

func tidy(root string) error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = root
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func cloneTSC(dir string) error {
	if err := git("", "clone",
		"--filter=blob:none",
		"--sparse",
		"--depth=1",
		"--branch=main",
		"--single-branch",
		upstreamRepo,
		dir,
	); err != nil {
		return err
	}
	// Cone mode keeps files directly in tsc/ (go.mod, LICENSE) and tsc/internal,
	// and does not fetch tsc/testdata or tsc/cmd.
	return git(dir, "sparse-checkout", "set", "tsc/internal")
}

func readGitRev(dir string) (rev, error) {
	r := rev{}
	for _, key := range revKeys {
		out, err := gitOutput(dir, "rev-parse", "HEAD:tsc/"+key)
		if err != nil {
			return nil, err
		}
		sha := strings.TrimSpace(out)
		if len(sha) != 40 {
			return nil, fmt.Errorf("unexpected sha for tsc/%s: %q", key, sha)
		}
		r[key] = sha
	}
	return r, nil
}

type contentEntry struct {
	Name string `json:"name"`
	SHA  string `json:"sha"`
}

func fetchRemoteRev() (rev, error) {
	req, err := http.NewRequest(http.MethodGet, contentsURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "startracex-tsgo-sync")
	if tok := os.Getenv("GITHUB_TOKEN"); tok != "" {
		req.Header.Set("Authorization", "Bearer "+tok)
	} else if tok := os.Getenv("GH_TOKEN"); tok != "" {
		req.Header.Set("Authorization", "Bearer "+tok)
	}
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github contents %s: %s", resp.Status, bytes.TrimSpace(body))
	}
	var entries []contentEntry
	if err := json.Unmarshal(body, &entries); err != nil {
		return nil, fmt.Errorf("parse tsc listing: %w", err)
	}
	r := rev{}
	for _, e := range entries {
		for _, key := range revKeys {
			if e.Name == key {
				r[key] = e.SHA
			}
		}
	}
	for _, key := range revKeys {
		if len(r[key]) != 40 {
			return nil, fmt.Errorf("tsc listing missing %s", key)
		}
	}
	return r, nil
}

func git(dir string, args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func gitOutput(dir string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}
	out, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	_, copyErr := io.Copy(out, in)
	closeErr := out.Close()
	if copyErr != nil {
		return copyErr
	}
	return closeErr
}

func setChanged(changed bool) error {
	path := os.Getenv("GITHUB_OUTPUT")
	if path == "" {
		return nil
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = fmt.Fprintf(f, "changed=%t\n", changed)
	return err
}
