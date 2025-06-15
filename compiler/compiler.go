package compiler

import (
	"fmt"

	"expression_parser/parser"
)

func NewCompiler() *Compiler {
	return &Compiler{program: NewProgram()}
}

type Compiler struct {
	program *Program
}

func (c *Compiler) Compile(node *parser.Node) (*Program, error) {
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

	if node.Type == parser.TypeOperation && (*node.Value.StringVal) == "VAR" {
		if err := c.doCompile(node.Params[1]); err != nil {
			return err
		}

		c.program.NewVariable(*node.Params[0].Value)
		return nil
	}

	if node.Type == parser.TypeOperation && (*node.Value.StringVal) == "IF" {
		if err := c.doCompile(node.Params[0]); err != nil {
			return err
		}

		c.program.NewIf(*node.Params[1].Value.StringVal, *node.Params[2].Value.StringVal)
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
