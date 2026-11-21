package fourslash_test

import (
	"testing"

	"github.com/startracex/tsgo/pkg/fourslash"
	"github.com/startracex/tsgo/pkg/testutil"
)

func TestGoToDefinitionShadowVariable(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `var shadowVariable = "foo";
function shadowVariableTestModule() {
    var /*shadowVariableDefinition*/shadowVariable;
    /*shadowVariableReference*/shadowVariable = 1;
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineGoToDefinition(t, false, "shadowVariableReference")
}
