package fourslash_test

import (
	"testing"

	"github.com/startracex/tsgo/pkg/fourslash"
	. "github.com/startracex/tsgo/pkg/fourslash/tests/util"
	"github.com/startracex/tsgo/pkg/ls/lsutil"
	"github.com/startracex/tsgo/pkg/testutil"
)

func TestCompletionImportModuleSpecifierEndingUnsupportedExtension(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `//@Filename:index.css
 body {}
//@Filename:module.ts
import ".//**/"`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete:	false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters:	&[]string{},
			EditRange:		Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Excludes: []string{
				"index.css",
			},
		},
		UserPreferences:	&lsutil.UserPreferences{ImportModuleSpecifierEnding: "js"},
	})
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete:	false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters:	&[]string{},
			EditRange:		Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Excludes: []string{
				"index",
			},
		},
		UserPreferences:	&lsutil.UserPreferences{ImportModuleSpecifierEnding: "index"},
	})
}
