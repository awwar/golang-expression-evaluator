package operation

import (
	"fmt"

	"expression_parser/compiler"
	"expression_parser/parser"
	"expression_parser/program"
	"expression_parser/tokenizer"
	"expression_parser/utility"
	"expression_parser/virtual_machine/operation"
)

func AppendProcedure(name string, op Procedure) {
	parser.AddProcedureParser(name, op)
	operation.AddExternalMethod(name, op)
	compiler.AddOperationSubCompiler(name, op)
}

func AppendExpression(name string, op Expression) {
	operation.AddExternalMethod(name, op)
}

type Procedure interface {
	Parse(token *tokenizer.Token, pr *parser.Parser) (*parser.Node, error)
	Execute(argc int, stack *utility.Stack[program.Value]) error
	Compile(program *program.Program, node *parser.Node, subcompile compiler.Subcompiler) error
}

type Expression interface {
	Execute(argc int, stack *utility.Stack[program.Value]) error
}

type CommonProcedure struct {
}

func (c *CommonProcedure) Parse(t *tokenizer.Token, pr *parser.Parser) (*parser.Node, error) {
	return parser.CreateAsOperation(t.Value, make([]*parser.Node, 2), t.Position), nil
}

func (c *CommonProcedure) Execute(_ int, _ *utility.Stack[program.Value]) error {
	return nil
}

func (c *CommonProcedure) Compile(_ *program.Program, node *parser.Node, cmp compiler.Subcompiler) error {
	return cmp.SubCompile(node)
}

type CommonExpression struct {
}

func (c *CommonExpression) Execute(_ int, _ *utility.Stack[program.Value]) error {
	return fmt.Errorf("implement me")
}
