package node

import (
	"fmt"

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
	// IF (rand() > 1) (#MORE, #LESS)
	expression, err := pr.SubparseOneInBracers()
	if err != nil {
		return nil, err
	}

	hashLinks, err := pr.SubparseListInBracers(2)
	if err != nil {
		return nil, err
	}

	if !hashLinks[0].IsFlowLink() || !hashLinks[1].IsFlowLink() {
		return nil, fmt.Errorf("if must have a 2 flow link")
	}

	return parser.CreateAsOperation(token.Value, []*parser.Node{expression, hashLinks[0], hashLinks[1]}, token.Position), nil
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
