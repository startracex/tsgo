package printer

import (
	"github.com/startracex/tsgo/pkg/ast"
	"github.com/startracex/tsgo/pkg/core"
	"github.com/startracex/tsgo/pkg/tspath"
)

// NOTE: EmitHost operations must be thread-safe
type EmitHost interface {
	Options() *core.CompilerOptions
	SourceFiles() []*ast.SourceFile
	ContentMapperExtensionRewrites() []tspath.ExtensionRewrite
	ContentMapperExtensionForFileName(fileName tspath.RootedFilePath) string
	CaseSensitivity() tspath.CaseSensitivity
	CommonSourceDirectory() tspath.RootedDirectoryPath
	IsEmitBlocked(file tspath.RootedFilePath) bool
	WriteFile(fileName tspath.RootedFilePath, text string) error
	GetEmitModuleFormatOfFile(file ast.HasFileName) core.ModuleKind
	NewEmitResolver(emitContext *EmitContext) EmitResolver
	IsSourceFileFromExternalLibrary(file *ast.SourceFile) bool
}
