package main

import (
	"fmt"
	"os"
	"strings"
)

// rev records the Git object SHAs of the upstream inputs we publish.
type rev map[string]string

func (r rev) equal(o rev) bool {
	if r == nil || o == nil {
		return false
	}
	for _, key := range revKeys {
		if r[key] == "" || r[key] != o[key] {
			return false
		}
	}
	return true
}

func formatRev(r rev) string {
	var b strings.Builder
	b.WriteString("# microsoft/TypeScript inputs published into pkg/.\n")
	b.WriteString("# testdata and cmd are not part of the sync.\n")
	for _, key := range revKeys {
		fmt.Fprintf(&b, "%s %s\n", key, r[key])
	}
	return b.String()
}

func parseRev(s string) (rev, error) {
	r := rev{}
	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		name, sha, ok := strings.Cut(line, " ")
		if !ok || !validSHA(sha) {
			return nil, fmt.Errorf("bad rev line %q", line)
		}
		r[name] = sha
	}
	for _, key := range revKeys {
		if !validSHA(r[key]) {
			return nil, fmt.Errorf("rev missing %s", key)
		}
	}
	return r, nil
}

func validSHA(s string) bool {
	if len(s) != 40 {
		return false
	}
	for _, c := range s {
		switch {
		case c >= '0' && c <= '9', c >= 'a' && c <= 'f':
		default:
			return false
		}
	}
	return true
}

func readRevFile(path string) (rev, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return parseRev(string(b))
}

func writeRevFile(path string, r rev) error {
	return os.WriteFile(path, []byte(formatRev(r)), 0o644)
}
