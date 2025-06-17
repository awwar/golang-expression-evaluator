package node

import (
	"expression_parser/compiler"
	"expression_parser/operation"
	"expression_parser/parser"
	"expression_parser/program"
	"expression_parser/tokenizer"
)

func init() {
	operation.AppendProcedure("IF", &If{})
}

type If struct {
	operation.CommonProcedure
}

func (i *If) Parse(token *tokenizer.Token, pr *parser.Parser) (*parser.Node, error) {
	expression, err := pr.SubparseOneInBracers()
	if err != nil {
		return nil, err
	}

	trueHashLink, err := pr.SubparseFlowLink()
	if err != nil {
		return nil, err
	}

	falseHashLinks, err := pr.SubparseFlowLink()
	if err != nil {
		return nil, err
	}

	return parser.CreateAsOperation(token.Value, []*parser.Node{expression, trueHashLink, falseHashLinks}, token.Position), nil
}

func (i *If) Compile(program *program.Program, node *parser.Node, subcompile compiler.Subcompiler) error {
	if err := subcompile(node.Params[0]); err != nil {
		return err
	}

	program.NewCSKP(2)
	program.NewJMP(*node.Params[2].Value.StringVal)
	program.NewSKP(1)
	program.NewJMP(*node.Params[1].Value.StringVal)

	return nil
}
