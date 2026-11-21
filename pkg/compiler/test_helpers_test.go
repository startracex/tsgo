package compiler_test

import (
	"github.com/startracex/tsgo/pkg/core"
	"github.com/startracex/tsgo/pkg/tspath"
)

func testFileNames(fileNames ...string) []tspath.RootedFilePath {
	return core.Map(fileNames, tspath.RootedFilePathFromNormalized)
}
