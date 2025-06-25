package compiler

import (
	"fmt"

	"expression_parser/parser"
	"expression_parser/program"
)

type Subcompiler interface {
	SubCompile(node *parser.Node) error
	GetMetadata(key string) string
}

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
	program  *program.Program
	metadata map[string]string
}

func (c *Compiler) GetMetadata(key string) string {
	return c.metadata[key]
}

func (c *Compiler) Compile(node *parser.Node) (*program.Program, error) {
	if err := c.doCompile(node); err != nil {
		return nil, err
	}

	return c.program, nil
}

func (c *Compiler) doCompile(node *parser.Node) error {
	if node == nil {
		return fmt.Errorf("compiler.Compile: nil node")
	}

	if node.Type == parser.TypeOperation {
		if proc, ok := OperationSubCompilerMap[node.Value.String()]; ok {
			return proc.Compile(c.program, node, c)
		}
	}

	return c.SubCompile(node)
}

func (c *Compiler) SubCompile(node *parser.Node) error {
	if node.Type == parser.TypeFlowDeclaration {
		returnType := ""
		bodyIndex := 0

		c.program.NewMark(node.Value.String())

		for i, n := range node.Params {
			bodyIndex = i

			if n.Type == parser.TypeConstant {
				returnType = n.Value.String()

				break
			}

			c.program.NewCall(n.Value.String(), len(n.Params))
			if len(n.Params) == 1 {
				c.program.NewVariable(n.Params[0].Value.String())
			}
		}

		if returnType == "" {
			return fmt.Errorf("flow declaration must contain response")
		}

		node.Params = node.Params[bodyIndex+1:]

		if err := checkLastOperationForReturn(node); err != nil {
			return err
		}
	}

	if node.Type == parser.TypeFlowBranchesDeclaration {
		c.program.NewMark(node.Value.String())
	}

	for _, child := range node.Params {
		if err := c.doCompile(child); err != nil {
			return err
		}
	}

	if node.Type == parser.TypeVariable {
		c.program.NewPush(node.Value)
	}

	if node.Type == parser.TypeOperation {
		c.program.NewCall(node.Value.String(), len(node.Params))
	}

	if node.Type == parser.TypeConstant {
		c.program.NewPush(node.Value)
	}

	return nil
}

func checkLastOperationForReturn(node *parser.Node) error {
	if len(node.Params) == 0 {
		return fmt.Errorf("flow must contain at least one operation")
	}

	lastBodyNode := node.Params[len(node.Params)-1]

	if lastBodyNode.Type != parser.TypeOperation || lastBodyNode.Value.String() != "RETURN" {
		return fmt.Errorf("flow declaration must contain RETURN")
	}

	return nil
}
