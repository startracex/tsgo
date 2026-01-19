package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/fourslash"
	"github.com/microsoft/typescript-go/testutil"
)

func TestAddSignaturePartial(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = ``
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.Insert(t, "interface Number { toFixed")
	f.Insert(t, "(")
}
