package fourslash_test

import (
	"testing"

	"github.com/startracex/tsgo/pkg/core"
	"github.com/startracex/tsgo/pkg/fourslash"
	"github.com/startracex/tsgo/pkg/ls/lsutil"
	"github.com/startracex/tsgo/pkg/testutil"
)

func TestRenameExportSpecifier(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: a.ts
const name = {};
export { name as name/**/ };
// @Filename: b.ts
import { name } from './a';
const x = name.toString();`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineRename(t, &lsutil.UserPreferences{ProvidePrefixAndSuffixTextForRename: core.TSFalse}, "")
}
