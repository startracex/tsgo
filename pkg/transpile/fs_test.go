package transpile

import (
	"testing"

	"github.com/startracex/tsgo/pkg/testutil"
	"github.com/startracex/tsgo/pkg/tspath"
)

func TestTranspileFSRejectsDirectoryAccess(t *testing.T) {
	t.Parallel()

	fs := &transpileFS{files: map[tspath.RootedFilePath]string{"/src/module.ts": ""}}
	testutil.AssertPanics(t, func() {
		fs.DirectoryExists(tspath.RootedDirectoryPath("/src"))
	}, `unexpected directory existence check for "/src"`)
	testutil.AssertPanics(t, func() {
		fs.Realpath(tspath.RootedFilePath("/src/module.ts").AsPath())
	}, `unexpected realpath request for "/src/module.ts"`)
}
