package fourslash_test

import (
	"testing"

	"github.com/startracex/tsgo/pkg/fourslash"
	"github.com/startracex/tsgo/pkg/testutil"
)

func TestTsxIncremental(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/**/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "")
	f.Insert(t, "<")
	f.Insert(t, "div")
	f.Insert(t, " ")
	f.Insert(t, " id")
	f.Insert(t, "=")
	f.Insert(t, "\"foo")
	f.Insert(t, "\"")
	f.Insert(t, ">")
}
