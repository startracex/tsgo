package fourslash_test

import (
	"testing"

	"github.com/startracex/tsgo/pkg/core"
	"github.com/startracex/tsgo/pkg/fourslash"
	"github.com/startracex/tsgo/pkg/ls/lsutil"
	"github.com/startracex/tsgo/pkg/lsp/lsproto"
	"github.com/startracex/tsgo/pkg/testutil"
)

func TestNavto_excludeLib1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @filename: /index.ts
import { weirdName as otherName } from "bar";
const [|weirdName|]: number = 1;
// @filename: /tsconfig.json
{}
// @filename: /node_modules/bar/index.d.ts
export const [|weirdName|];
// @filename: /node_modules/bar/package.json
{}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyWorkspaceSymbol(t, []*fourslash.VerifyWorkspaceSymbolCase{
		{
			Pattern:	"weirdName",
			Preferences:	&lsutil.UserPreferences{ExcludeLibrarySymbolsInNavTo: core.TSFalse},
			Exact: new([]*lsproto.SymbolInformation{
				{
					Name:		"weirdName",
					Kind:		lsproto.SymbolKindVariable,
					Location:	f.Ranges()[0].LSLocation(),
				},
				{
					Name:		"weirdName",
					Kind:		lsproto.SymbolKindVariable,
					Location:	f.Ranges()[1].LSLocation(),
				},
			}),
		},
	})
	f.VerifyWorkspaceSymbol(t, []*fourslash.VerifyWorkspaceSymbolCase{
		{
			Pattern:	"weirdName",
			Preferences:	nil,
			Exact: new([]*lsproto.SymbolInformation{
				{
					Name:		"weirdName",
					Kind:		lsproto.SymbolKindVariable,
					Location:	f.Ranges()[0].LSLocation(),
				},
			}),
		},
	})
}
