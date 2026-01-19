package main

import (
	"os"

	"github.com/startracex/go-pub"
)

func main() {
	outDir := "pkg/tsgo"
	err := os.RemoveAll(outDir)
	if err != nil {
		panic(err)
	}
	r, err := pub.NewModRoot("submodules/typescript-go/go.mod")
	if err != nil {
		panic(err)
	}
	pc := &pub.PublicConfig{
		OutDir:         outDir,
		InternalOutDir: outDir,
		IncludeModFile: true,
		IncludeSumFile: true,
	}
	err = pc.Public(r)
	if err != nil {
		panic(err)
	}
}
