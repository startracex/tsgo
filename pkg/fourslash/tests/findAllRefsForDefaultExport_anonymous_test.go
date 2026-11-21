package fourslash_test

import (
	"testing"

	"github.com/startracex/tsgo/pkg/fourslash"
	"github.com/startracex/tsgo/pkg/testutil"
)

func TestFindAllRefsForDefaultExport_anonymous(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.ts
export /*1*/default 1;
// @Filename: /b.ts
import a from "./a";`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineFindAllReferences(t, "1")
}

func TestFindAllRefsForDefaultExport_anonymousClass(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.ts
/*export*/export /*default*/default /*class*/class {}
// @Filename: /b.ts
import /*import*/AnonymousClass from "./a";
const instance = new /*use*/AnonymousClass();`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineFindAllReferences(t, "export", "default", "class", "import", "use")
}

func TestFindAllRefsForDefaultExport_anonymousFunction(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.ts
/*export*/export /*default*/default /*async*/async /*function*/function () {}
// @Filename: /b.ts
import /*import*/anonymousFunction from "./a";
/*use*/anonymousFunction();`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineFindAllReferences(t, "export", "default", "async", "function", "import", "use")
}
