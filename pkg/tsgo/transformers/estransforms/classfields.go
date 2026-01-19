package estransforms

import (
	"github.com/microsoft/typescript-go/ast"
	"github.com/microsoft/typescript-go/transformers"
)

type classFieldsTransformer struct {
	transformers.Transformer
}

func (ch *classFieldsTransformer) visit(node *ast.Node) *ast.Node {
	return node	// !!!
}

func newClassFieldsTransformer(opts *transformers.TransformOptions) *transformers.Transformer {
	tx := &classFieldsTransformer{}
	return tx.NewTransformer(tx.visit, opts.Context)
}
