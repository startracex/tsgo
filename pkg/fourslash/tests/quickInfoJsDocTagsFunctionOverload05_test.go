package fourslash_test

import (
	"testing"

	"github.com/startracex/tsgo/pkg/fourslash"
	"github.com/startracex/tsgo/pkg/testutil"
)

func TestQuickInfoJsDocTagsFunctionOverload05(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: quickInfoJsDocTagsFunctionOverload05.ts
declare function /*1*/foo(): void;

/**
 * @tag Tag text
 */
declare function /*2*/foo(x: number): void`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineHover(t)
}
