package fourslash_test

import (
	"testing"

	"github.com/startracex/tsgo/pkg/fourslash"
	"github.com/startracex/tsgo/pkg/testutil"
)

func TestThisBindingInLambda(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class Greeter {
    constructor() { 
		[].forEach((anything)=>{
			console.log(th/**/is);
		});
	}
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyQuickInfoAt(t, "", "this: this", "")
}
