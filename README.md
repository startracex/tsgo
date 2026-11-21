# tsgo

Import TypeScript's Go compiler packages from this module. Sources are synced from [`microsoft/TypeScript` `tsc`](https://github.com/microsoft/TypeScript/tree/main/tsc); `internal` packages are republished under `pkg/`.

```bash
go get github.com/startracex/tsgo
```

Requires Go 1.27.

```go
package main

import (
	"fmt"

	"github.com/startracex/tsgo/pkg/ast"
	"github.com/startracex/tsgo/pkg/core"
	"github.com/startracex/tsgo/pkg/parser"
	"github.com/startracex/tsgo/pkg/tspath"
)

func main() {
	fileName := tspath.RootedFilePathFromNormalized("/hello.ts")
	file := parser.ParseSourceFile(ast.SourceFileParseOptions{
		FileName: fileName,
		PathKey:  tspath.CaseSensitive.PathKey(tspath.RootedPath(fileName)),
	}, "export const answer = 42", core.ScriptKindTS)
	fmt.Println(len(file.Statements.Nodes))
}
```

Published sources keep the [upstream license](LICENSE).
