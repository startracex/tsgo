package fourslash_test

import (
	"testing"

	"github.com/startracex/tsgo/pkg/fourslash"
	"github.com/startracex/tsgo/pkg/testutil"
)

func TestAddMemberToModule(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `namespace A {
    /*var*/
}
module /*check*/A {
    var p;
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "check")
	f.VerifyQuickInfoExists(t)
	f.GoToMarker(t, "var")
	f.Insert(t, "var o;")
	f.GoToMarker(t, "check")
	f.VerifyQuickInfoExists(t)
}
