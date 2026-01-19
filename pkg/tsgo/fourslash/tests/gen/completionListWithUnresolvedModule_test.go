package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/fourslash"
	. "github.com/microsoft/typescript-go/fourslash/tests/util"
	"github.com/microsoft/typescript-go/ls"
	"github.com/microsoft/typescript-go/lsp/lsproto"
	"github.com/microsoft/typescript-go/testutil"
)

func TestCompletionListWithUnresolvedModule(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `module m {
    import foo = module('_foo');
    var n: num/**/
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
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:		"number",
					SortText:	PtrTo(string(ls.SortTextGlobalsOrKeywords)),
				},
			},
		},
	})
}
