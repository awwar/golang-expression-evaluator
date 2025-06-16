package procedure

import (
	"expression_parser/parser"
	"expression_parser/program"
)

func init() {
	AppendProcedure("IF", &If{})
}

type If struct {
	CommonProcedure
}

func (i *If) Compile(program *program.Program, node *parser.Node, subcompile func(node *parser.Node) error) error {
	if err := subcompile(node.Params[0]); err != nil {
		return err
	}

	program.NewCSKP(2)
	program.NewJMP(*node.Params[2].Value.StringVal)
	program.NewSKP(1)
	program.NewJMP(*node.Params[1].Value.StringVal)

	return nil
}
