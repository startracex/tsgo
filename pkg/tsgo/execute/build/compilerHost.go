package build

import (
	"github.com/microsoft/typescript-go/ast"
	"github.com/microsoft/typescript-go/compiler"
	"github.com/microsoft/typescript-go/diagnostics"
	"github.com/microsoft/typescript-go/tsoptions"
	"github.com/microsoft/typescript-go/tspath"
	"github.com/microsoft/typescript-go/vfs"
)

type compilerHost struct {
	host	*host
	trace	func(msg *diagnostics.Message, args ...any)
}

var _ compiler.CompilerHost = (*compilerHost)(nil)

func (h *compilerHost) FS() vfs.FS {
	return h.host.FS()
}

func (h *compilerHost) DefaultLibraryPath() string {
	return h.host.DefaultLibraryPath()
}

func (h *compilerHost) GetCurrentDirectory() string {
	return h.host.GetCurrentDirectory()
}

func (h *compilerHost) Trace(msg *diagnostics.Message, args ...any) {
	h.trace(msg, args...)
}

func (h *compilerHost) GetSourceFile(opts ast.SourceFileParseOptions) *ast.SourceFile {
	return h.host.GetSourceFile(opts)
}

func (h *compilerHost) GetResolvedProjectReference(fileName string, path tspath.Path) *tsoptions.ParsedCommandLine {
	return h.host.GetResolvedProjectReference(fileName, path)
}
