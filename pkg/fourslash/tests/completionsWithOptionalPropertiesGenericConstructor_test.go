package fourslash_test

import (
	"testing"

	"github.com/startracex/tsgo/pkg/fourslash"
	. "github.com/startracex/tsgo/pkg/fourslash/tests/util"
	"github.com/startracex/tsgo/pkg/ls"
	"github.com/startracex/tsgo/pkg/lsp/lsproto"
	"github.com/startracex/tsgo/pkg/testutil"
)

func TestCompletionsWithOptionalPropertiesGenericConstructor(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strict: true
interface Options {
    someFunction?: () => string
    anotherFunction?: () => string
}

export class Clazz<T extends Options> {
    constructor(public a: T) {}
}

new Clazz({ /*1*/ })`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		IsIncomplete:	false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters:	&DefaultCommitCharacters,
			EditRange:		Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:		"someFunction?",
					InsertText:	new("someFunction"),
					FilterText:	new("someFunction"),
					SortText:	new(string(ls.SortTextOptionalMember)),
				},
				&lsproto.CompletionItem{
					Label:		"anotherFunction?",
					InsertText:	new("anotherFunction"),
					FilterText:	new("anotherFunction"),
					SortText:	new(string(ls.SortTextOptionalMember)),
				},
			},
		},
	})
}
