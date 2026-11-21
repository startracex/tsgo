package main

import (
	"bytes"
	"strings"
	"testing"
)

func sampleRev() rev {
	return rev{
		"internal": strings.Repeat("a", 40),
		"go.mod":   strings.Repeat("b", 40),
		"go.sum":   strings.Repeat("c", 40),
		"LICENSE":  strings.Repeat("d", 40),
	}
}

func TestRevRoundTrip(t *testing.T) {
	got, err := parseRev(formatRev(sampleRev()))
	if err != nil {
		t.Fatal(err)
	}
	if !sampleRev().equal(got) {
		t.Fatalf("round trip = %#v", got)
	}
}

func TestRevMismatch(t *testing.T) {
	other := sampleRev()
	other["go.mod"] = strings.Repeat("e", 40)
	if sampleRev().equal(other) {
		t.Fatal("different go.mod SHAs compared equal")
	}
}

func TestRewriteImportPrefixLeavesLiterals(t *testing.T) {
	src := []byte("package p\n\nimport (\n\t\"github.com/microsoft/TypeScript/tsc/ast\"\n\t\"github.com/microsoft/TypeScript/tsc/vfs/internal\"\n)\n\nconst stack = `github.com/microsoft/TypeScript/tsc/internal/lsp`\nconst gen = \"\\t\\\"github.com/microsoft/TypeScript/tsc/internal/json\\\"\\n\"\n// see https://github.com/microsoft/TypeScript/tsc/issues/2953\n")
	got, err := rewriteImportBytes(src)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		`"github.com/startracex/tsgo/pkg/ast"`,
		`"github.com/startracex/tsgo/pkg/vfs/internal"`,
		`github.com/startracex/tsgo/pkg/json`,
		"`github.com/microsoft/TypeScript/tsc/internal/lsp`",
		"https://github.com/microsoft/TypeScript/tsc/issues/2953",
	} {
		if !bytes.Contains(got, []byte(want)) {
			t.Fatalf("missing %s in %s", want, got)
		}
	}
	if bytes.Contains(got, []byte("/pkg/internal/")) {
		t.Fatalf("module-level internal survived: %s", got)
	}
}
