package procedure

import (
	"expression_parser/parser"
	"expression_parser/program"
	"expression_parser/utility"
)

var (
	Map = map[string]Procedure{}
)

func AppendProcedure(name string, op Procedure) {
	Map[name] = op
}

type Procedure interface {
	Execute(argc int, stack *utility.Stack[program.Value]) error
	Compile(program *program.Program, node *parser.Node, subcompile func(node *parser.Node) error) error
}

type CommonProcedure struct {
}

func (c *CommonProcedure) Execute(argc int, stack *utility.Stack[program.Value]) error {
	return nil
}

func (c *CommonProcedure) Compile(_ *program.Program, node *parser.Node, subcompile func(node *parser.Node) error) error {
	return subcompile(node)
}
