package fourslash_test

import (
	"testing"

	"github.com/startracex/tsgo/pkg/fourslash"
	. "github.com/startracex/tsgo/pkg/fourslash/tests/util"
	"github.com/startracex/tsgo/pkg/ls/lsutil"
	"github.com/startracex/tsgo/pkg/lsp/lsproto"
	"github.com/startracex/tsgo/pkg/testutil"
)

func TestJsxTagNameCompletionWithExistingJsxInitializer(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @filename: /foo.tsx
declare namespace JSX {
    interface Element { }
    interface IntrinsicElements {
        foo: {
            className: string;
        }
    }
}
<foo cl/**/={""} />`
	f, done := fourslash.NewFourslash(t, fourslash.GetDefaultCapabilitiesWithOptions(&fourslash.ClientCapabilitiesOptions{
		CompletionItem: &lsproto.ClientCompletionItemOptions{
			SnippetSupport: new(true),
		},
	}), content)
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
					Label:	"className",
					Detail:	new("(property) className: string"),
				},
			},
		},
		UserPreferences:	&lsutil.UserPreferences{JsxAttributeCompletionStyle: lsutil.JsxAttributeCompletionStyleBraces},
	})
}
