package printer

import (
	"github.com/startracex/tsgo/pkg/ast"
	"github.com/startracex/tsgo/pkg/tspath"
)

type SourceFileMetaDataProvider interface {
	GetSourceFileMetaData(path tspath.PathKey) *ast.SourceFileMetaData
}
