package node

import (
	"expression_parser/compiler"
	"expression_parser/operation"
	"expression_parser/parser"
	"expression_parser/program"
	"expression_parser/tokenizer"
)

func init() {
	operation.AppendProcedure("VAR", &Var{})
}

type Var struct {
	operation.CommonProcedure
}

func (i *Var) Parse(token *tokenizer.Token, pr *parser.Parser) (*parser.Node, error) {
	expression, err := pr.SubparseOneInBracers()
	if err != nil {
		return nil, err
	}

	variableNode, err := pr.SubparseVariableName()
	if err != nil {
		return nil, err
	}

	return parser.CreateAsOperation(token.Value, []*parser.Node{variableNode, expression}, token.Position), nil
}

func (i *Var) Compile(program *program.Program, node *parser.Node, cmp compiler.Subcompiler) error {
	if err := cmp.SubCompile(node.Params[1]); err != nil {
		return err
	}

	program.NewVariable(node.Params[0].Value.String())

	return nil
}
