package sourcemap

import (
	"github.com/startracex/tsgo/pkg/core"
	"github.com/startracex/tsgo/pkg/tspath"
)

type Source interface {
	Text() string
	FileName() tspath.RootedFilePath
	ECMALineMap() []core.TextPos
}
