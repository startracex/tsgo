package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/fourslash"
	. "github.com/microsoft/typescript-go/fourslash/tests/util"
	"github.com/microsoft/typescript-go/ls"
	"github.com/microsoft/typescript-go/lsp/lsproto"
	"github.com/microsoft/typescript-go/testutil"
)

func TestCompletionsOverridingMethod4(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @newline: LF
// @Filename: secret.ts
class Secret {
    #secret(): string {
        return "secret";
    }

    private tell(): string {
        return this.#secret();
    }

    protected hint(): string {
        return "hint";
    }

    public refuse(): string {
        return "no comments";
    }
}

class Gossip extends Secret {
    /* no telling secrets */
    /*a*/
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "a", &fourslash.CompletionsExpectedList{
		IsIncomplete:	false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters:	&[]string{},
			EditRange:		Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:		"hint",
					InsertText:	PtrTo("protected hint(): string {\n}"),
					FilterText:	PtrTo("hint"),
					SortText:	PtrTo(string(ls.SortTextLocationPriority)),
				},
				&lsproto.CompletionItem{
					Label:		"refuse",
					InsertText:	PtrTo("public refuse(): string {\n}"),
					FilterText:	PtrTo("refuse"),
					SortText:	PtrTo(string(ls.SortTextLocationPriority)),
				},
			},
			Excludes: []string{
				"tell",
				"#secret",
			},
		},
	})
}
