package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/fourslash"
	. "github.com/microsoft/typescript-go/fourslash/tests/util"
	"github.com/microsoft/typescript-go/lsp/lsproto"
	"github.com/microsoft/typescript-go/testutil"
)

func TestImportMetaCompletionDetails(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @filename: index.mts
// @module: Node16
// @strict: true
let x = import.meta/**/;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete:	false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters:	&DefaultCommitCharacters,
			EditRange:		Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:	"meta",
					Detail:	PtrTo("(property) ImportMetaExpression.meta: ImportMeta"),
				},
			},
		},
	})
	f.VerifyNoErrors(t)
}
