package operation

import (
	"fmt"
	"strings"

	"expression_parser/compiler"
	"expression_parser/parser"
	"expression_parser/program"
	"expression_parser/tokenizer"
	"expression_parser/utility"
	"expression_parser/virtual_machine/operation"
)

func AppendProcedure(name string, op Procedure) {
	n := strings.ToLower(name)

	parser.AddProcedureParser(n, op)
	operation.AddExternalMethod(n, op)
	compiler.AddOperationSubCompiler(n, op)
}

func AppendExpression(name string, op Expression) {
	n := strings.ToLower(name)

	operation.AddExternalMethod(n, op)
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
