package node

import (
	"expression_parser/compiler"
	"expression_parser/operation"
	"expression_parser/parser"
	"expression_parser/program"
	"expression_parser/tokenizer"
)

func init() {
	operation.AppendProcedure("RETURN", &Return{})
}

type Return struct {
	operation.CommonProcedure
}

func (i *Return) Parse(token *tokenizer.Token, pr *parser.Parser) (*parser.Node, error) {
	// RETURN (expr)
	expr, err := pr.SubparseOneInBracers()
	if err != nil {
		return nil, err
	}

	return parser.CreateAsOperation(token.Value, []*parser.Node{expr}, token.Position), nil
}

func (i *Return) Compile(_ *program.Program, node *parser.Node, subcompile compiler.Subcompiler) error {
	return subcompile(node.Params[0])
}
