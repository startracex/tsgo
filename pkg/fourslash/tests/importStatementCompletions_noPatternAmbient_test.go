package fourslash_test

import (
	"testing"

	"github.com/startracex/tsgo/pkg/fourslash"
	. "github.com/startracex/tsgo/pkg/fourslash/tests/util"
	"github.com/startracex/tsgo/pkg/ls"
	"github.com/startracex/tsgo/pkg/lsp/lsproto"
	"github.com/startracex/tsgo/pkg/testutil"
)

func TestImportStatementCompletions_noPatternAmbient(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /types.d.ts
declare module "*.css" {
  const styles: any;
  export = styles;
}
// @Filename: /index.ts
import style/**/`
	f, done := fourslash.NewFourslash(t, fourslash.GetDefaultCapabilitiesWithOptions(&fourslash.ClientCapabilitiesOptions{
		CompletionItem: &lsproto.ClientCompletionItemOptions{
			SnippetSupport: new(true),
		},
	}), content)
	defer done()
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete:	false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters:	&[]string{},
			EditRange:		Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:		"type",
					SortText:	new(string(ls.SortTextGlobalsOrKeywords)),
				},
			},
		},
	})
}
