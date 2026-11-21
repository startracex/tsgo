package fourslash_test

import (
	"testing"

	"github.com/startracex/tsgo/pkg/fourslash"
	"github.com/startracex/tsgo/pkg/testutil"
)

func TestCallHierarchyIncomingCallsObjectLiteralMethodInStringLiteralComputedProperty(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `const obj = {
  ["x"]: {
    method() {
      return ""./*split*/split(",");
    }
  }
};
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "split")
	f.VerifyBaselineCallHierarchy(t)
}
