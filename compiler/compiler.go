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

	for _, child := range node.Params {
		if err := c.doCompile(child); err != nil {
			return err
		}
	}

	if node.Type == parser.TypeOperation {
		c.program.NewCall(*node.Value.StringVal, len(node.Params))
	}

	if node.Type == parser.TypeConstant {
		c.program.NewPush(*node.Value)
	}

	return nil
}
