package fourslash_test

import (
	"testing"

	"github.com/startracex/tsgo/pkg/core"
	"github.com/startracex/tsgo/pkg/fourslash"
	"github.com/startracex/tsgo/pkg/ls/lsutil"
	"github.com/startracex/tsgo/pkg/lsp/lsproto"
	"github.com/startracex/tsgo/pkg/testutil"
)

func TestOrganizeImportsUnicode2(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `import {
    a2,
    a100,
    a1,
} from './foo';

console.log(a1, a2, a100);`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`import {
    a1,
    a100,
    a2,
} from './foo';

console.log(a1, a2, a100);`,
		lsproto.CodeActionKindSourceOrganizeImportsTs,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase:		core.TSFalse,
			OrganizeImportsCollation:		lsutil.OrganizeImportsCollationUnicode,
			OrganizeImportsNumericCollation:	core.TSFalse,
		},
	)
	f.VerifyOrganizeImports(t,
		`import {
    a1,
    a2,
    a100,
} from './foo';

console.log(a1, a2, a100);`,
		lsproto.CodeActionKindSourceOrganizeImportsTs,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase:		core.TSFalse,
			OrganizeImportsCollation:		lsutil.OrganizeImportsCollationUnicode,
			OrganizeImportsNumericCollation:	core.TSTrue,
		},
	)
}
