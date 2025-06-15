package procedure

import (
	"expression_parser/parser"
	"expression_parser/program"
)

func init() {
	AppendProcedure("VAR", &Var{})
}

type Var struct {
	CommonProcedure
}

func (i *Var) Compile(program *program.Program, node *parser.Node, subcompile func(node *parser.Node) error) error {
	if err := subcompile(node.Params[1]); err != nil {
		return err
	}

	program.NewVariable(*node.Params[0].Value.StringVal)

	return nil
}
