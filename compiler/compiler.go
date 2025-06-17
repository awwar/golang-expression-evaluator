package compiler

import (
	"fmt"

	"expression_parser/parser"
	"expression_parser/program"
)

type Subcompiler func(node *parser.Node) error

type OperationSubCompiler interface {
	Compile(program *program.Program, node *parser.Node, subcompile Subcompiler) error
}

var OperationSubCompilerMap = map[string]OperationSubCompiler{}

func AddOperationSubCompiler(name string, subCompiler OperationSubCompiler) {
	OperationSubCompilerMap[name] = subCompiler
}

func NewCompiler() *Compiler {
	return &Compiler{program: program.NewProgram()}
}

type Compiler struct {
	program *program.Program
}

func (c *Compiler) Compile(node *parser.Node) (*program.Program, error) {
	err := c.doCompile(node)
	if err != nil {
		return nil, err
	}

	return c.program, nil
}

func (c *Compiler) doCompile(node *parser.Node) error {
	if node == nil {
		return fmt.Errorf("compiler.Compile: nil node")
	}

	if node.Type == parser.TypeOperation {
		if proc, ok := OperationSubCompilerMap[*node.Value.StringVal]; ok {
			return proc.Compile(c.program, node, c.subCompile)
		}
	}

	return c.subCompile(node)

}

func (c *Compiler) subCompile(node *parser.Node) error {
	if node.Type == parser.TypeFlowMetadata {
		return nil
	}

	if node.Type == parser.TypeFlowDeclaration {
		c.program.NewMark(*node.Value.StringVal)
	}

	for _, child := range node.Params {
		if err := c.doCompile(child); err != nil {
			return err
		}
	}

	if node.Type == parser.TypeVariable {
		c.program.NewPush(*node.Value)
	}

	if node.Type == parser.TypeOperation {
		c.program.NewCall(*node.Value.StringVal, len(node.Params))
	}

	if node.Type == parser.TypeConstant {
		c.program.NewPush(*node.Value)
	}

	return nil
}
