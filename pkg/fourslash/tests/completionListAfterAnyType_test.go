package fourslash_test

import (
	"testing"

	"github.com/startracex/tsgo/pkg/fourslash"
	. "github.com/startracex/tsgo/pkg/fourslash/tests/util"
	"github.com/startracex/tsgo/pkg/testutil"
)

func TestCompletionListAfterAnyType(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare class myString {
    charAt(pos: number): string;
}

function bar(a: myString) {
    var x: any = a./**/
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
				"charAt",
			},
		},
	})
}
