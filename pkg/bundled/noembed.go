//go:build noembed

package bundled

import (
	"fmt"
	"sync"
	"testing"

	"github.com/startracex/tsgo/pkg/osutil"
	"github.com/startracex/tsgo/pkg/tspath"
	"github.com/startracex/tsgo/pkg/vfs"
	"github.com/startracex/tsgo/pkg/vfs/osvfs"
)

const embedded = false

func wrapFS(fs vfs.FS) vfs.FS {
	return fs
}

var executableDir = sync.OnceValue(func() tspath.RootedDirectoryPath {
	exe, err := osutil.Executable()
	if err != nil {
		panic(fmt.Sprintf("bundled: failed to get executable path: %v", err))
	}
	realExe := osvfs.FS().Realpath(tspath.RootedFilePathFromAbsolute(exe).AsPath())
	return realExe.Directory()
})

var libPath = sync.OnceValue(func() string {
	if testing.Testing() {
		return TestingLibPath()
	}
	dir := executableDir()

	libdts := dir.ResolveFile("lib.d.ts")
	if info := osvfs.FS().Stat(libdts.AsPath()); info == nil {
		panic(fmt.Sprintf("bundled: %v does not exist; this executable may be misplaced", libdts))
	}

	return dir.AsString()
})

func IsBundled(path string) bool {
	return false
}
