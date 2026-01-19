package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/fourslash"
	. "github.com/microsoft/typescript-go/fourslash/tests/util"
	"github.com/microsoft/typescript-go/ls"
	"github.com/microsoft/typescript-go/lsp/lsproto"
	"github.com/microsoft/typescript-go/testutil"
)

func TestJavascriptModules21(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @module: system
// @Filename: mod.js
function foo() { return {a: true}; }
module.exports = foo();
// @Filename: app.js
import mod from "./mod"
mod./**/`
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
					Label:	"a",
					Kind:	PtrTo(lsproto.CompletionItemKindField),
				},
				&lsproto.CompletionItem{
					Label:		"mod",
					Kind:		PtrTo(lsproto.CompletionItemKindText),
					SortText:	PtrTo(string(ls.SortTextJavascriptIdentifiers)),
				},
			},
		},
	})
}
