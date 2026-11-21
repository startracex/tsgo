package ls

import (
	"github.com/startracex/tsgo/pkg/ls/autoimport"
	"github.com/startracex/tsgo/pkg/ls/lsconv"
	"github.com/startracex/tsgo/pkg/ls/lsutil"
	"github.com/startracex/tsgo/pkg/sourcemap"
	"github.com/startracex/tsgo/pkg/tspath"
)

type WorkspaceFolder struct {
	Path	tspath.RootedDirectoryPath
	Name	string
}

type Host interface {
	CaseSensitivity() tspath.CaseSensitivity
	ReadFile(path tspath.RootedFilePath) (contents string, ok bool)
	Converters() *lsconv.Converters
	GetPreferences(activeFile string) lsutil.UserPreferences
	GetWorkspaceFolders() []WorkspaceFolder
	GetECMALineInfo(fileName tspath.RootedFilePath) *sourcemap.ECMALineInfo
	AutoImportRegistry() *autoimport.Registry

	// Used for module specifier completions.
	// ! Do not use for anything else, as this violates the principle that
	// the host is a snapshot-in-time.
	ReadDirectory(path tspath.RootedDirectoryPath, extensions []string, excludes []string, includes []string, depth int) []tspath.RootedFilePath
	GetDirectories(path tspath.RootedDirectoryPath) []string
	DirectoryExists(path tspath.RootedDirectoryPath) bool
	FileExists(path tspath.RootedFilePath) bool
}
