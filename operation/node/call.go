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

	args, err := pr.SubparseListInBracers(-1)
	if err != nil {
		return nil, err
	}

	variable, err := pr.SubparseVariableName()
	if err != nil {
		return nil, err
	}

	subnodes := []*parser.Node{link, variable}

	subnodes = append(subnodes, args...)

	return parser.CreateAsOperation(token.Value, subnodes, token.Position), nil
}

func (i *Call) Compile(program *program.Program, node *parser.Node, cmp compiler.Subcompiler) error {
	for _, a := range node.Params[2:] {
		err := cmp.SubCompile(a)
		if err != nil {
			return err
		}
	}

	program.NewJMP(node.Params[0].Value.String())
	program.NewVariable(node.Params[1].Value.String())

	return nil
}
