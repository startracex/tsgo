package fourslash_test

import (
	"testing"

	"github.com/startracex/tsgo/pkg/fourslash"
	. "github.com/startracex/tsgo/pkg/fourslash/tests/util"
	"github.com/startracex/tsgo/pkg/ls"
	"github.com/startracex/tsgo/pkg/lsp/lsproto"
	"github.com/startracex/tsgo/pkg/testutil"
)

func TestCompletionsWithDeprecatedTag7(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strict: true
interface I {
    /** @deprecated a */
    a: number;
}
const foo = {
    a: 1
}
const i: I = {
    ...foo,
    /**/
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete:	false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters:	&DefaultCommitCharacters,
			EditRange:		Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:		"a",
					SortText:	new(string(ls.DeprecateSortText(ls.SortTextMemberDeclaredBySpreadAssignment))),
					Kind:		new(lsproto.CompletionItemKindField),
					Tags:		&[]lsproto.CompletionItemTag{lsproto.CompletionItemTagDeprecated},
				},
			},
		},
	})
}
