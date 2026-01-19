package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/fourslash"
	. "github.com/microsoft/typescript-go/fourslash/tests/util"
	"github.com/microsoft/typescript-go/ls"
	"github.com/microsoft/typescript-go/lsp/lsproto"
	"github.com/microsoft/typescript-go/testutil"
)

func TestCompletionFilterText4(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare const x: [number, number];
x[|.|]/**/;
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete:	false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:		"0",
					Kind:		PtrTo(lsproto.CompletionItemKindField),
					SortText:	PtrTo(string(ls.SortTextLocationPriority)),
					InsertText:	PtrTo("[0]"),
					FilterText:	PtrTo(".0"),
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						TextEdit: &lsproto.TextEdit{
							NewText:	"[0]",
							Range:		f.Ranges()[0].LSRange,
						},
					},
				},
			},
		},
	})
}
