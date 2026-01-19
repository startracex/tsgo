package ls

import (
	"github.com/microsoft/typescript-go/format"
	"github.com/microsoft/typescript-go/ls/autoimport"
	"github.com/microsoft/typescript-go/ls/lsconv"
	"github.com/microsoft/typescript-go/ls/lsutil"
	"github.com/microsoft/typescript-go/sourcemap"
)

type Host interface {
	UseCaseSensitiveFileNames() bool
	ReadFile(path string) (contents string, ok bool)
	Converters() *lsconv.Converters
	UserPreferences() *lsutil.UserPreferences
	FormatOptions() *format.FormatCodeSettings
	GetECMALineInfo(fileName string) *sourcemap.ECMALineInfo
	AutoImportRegistry() *autoimport.Registry
}
