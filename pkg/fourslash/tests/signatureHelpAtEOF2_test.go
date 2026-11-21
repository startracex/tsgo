package fourslash_test

import (
	"testing"

	"github.com/startracex/tsgo/pkg/fourslash"
	"github.com/startracex/tsgo/pkg/lsp/lsproto"
	"github.com/startracex/tsgo/pkg/testutil"
)

func TestSignatureHelpAtEOF2(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `console.log()
/**/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyNoSignatureHelpForMarkersWithContext(t, &lsproto.SignatureHelpContext{TriggerKind: lsproto.SignatureHelpTriggerKindInvoked}, "")
}
