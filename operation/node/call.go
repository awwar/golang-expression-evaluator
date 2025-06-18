package node

import (
	"expression_parser/compiler"
	"expression_parser/operation"
	"expression_parser/parser"
	"expression_parser/program"
	"expression_parser/tokenizer"
)

func init() {
	operation.AppendProcedure("CALL", &Call{})
}

type Call struct {
	operation.CommonProcedure
}

func (i *Call) Parse(token *tokenizer.Token, pr *parser.Parser) (*parser.Node, error) {
	// CALL #NAME $RESULT
	link, err := pr.SubparseFlowLink()
	if err != nil {
		return nil, err
	}

	variable, err := pr.SubparseVariableName()
	if err != nil {
		return nil, err
	}

	return parser.CreateAsOperation(token.Value, []*parser.Node{link, variable}, token.Position), nil
}

func (i *Call) Compile(program *program.Program, node *parser.Node, subcompile compiler.Subcompiler) error {
	program.NewJMP(*node.Params[0].Value.StringVal)
	program.NewVariable(*node.Params[1].Value.StringVal)

	return nil
}
