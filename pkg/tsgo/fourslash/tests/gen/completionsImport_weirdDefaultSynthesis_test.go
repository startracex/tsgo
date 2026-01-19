package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/fourslash"
	. "github.com/microsoft/typescript-go/fourslash/tests/util"
	"github.com/microsoft/typescript-go/testutil"
)

func TestCompletionsImport_weirdDefaultSynthesis(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @module: commonjs
// @esModuleInterop: false
// @allowSyntheticDefaultImports: false
// @Filename: /collection.ts
class Collection {
  public static readonly default: typeof Collection = Collection;
}
export = Collection as typeof Collection & { default: typeof Collection };
// @Filename: /index.ts
Colle/**/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyApplyCodeActionFromCompletion(t, PtrTo(""), &fourslash.ApplyCodeActionFromCompletionOptions{
		Name:		"Collection",
		Source:		"./collection",
		Description:	"Add import from \"./collection\"",
		NewFileContent: PtrTo(`import Collection = require("./collection");

Colle`),
	})
}
