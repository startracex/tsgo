package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/fourslash"
	. "github.com/microsoft/typescript-go/fourslash/tests/util"
	"github.com/microsoft/typescript-go/lsp/lsproto"
	"github.com/microsoft/typescript-go/testutil"
)

func TestNavigateToSymbolIterator(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class C {
    [|[Symbol.iterator]() {}|]
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyWorkspaceSymbol(t, []*fourslash.VerifyWorkspaceSymbolCase{
		{
			Pattern:	"iterator",
			Preferences:	nil,
			Exact: PtrTo([]*lsproto.SymbolInformation{
				{
					Name:		"iterator",
					Kind:		lsproto.SymbolKindMethod,
					Location:	f.Ranges()[0].LSLocation(),
					ContainerName:	PtrTo("C"),
				},
			}),
		},
	})
}
