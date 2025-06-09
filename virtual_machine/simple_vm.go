package virtual_machine

import (
	"fmt"

	"expression_parser/parser"
	"expression_parser/virtual_machine/operations"
)

var (
	operationsMap = map[string]Operation{
		"sum":        operations.Sum,
		"uppercase":  operations.Uppercase,
		"trim_space": operations.TrimSpace,
		"+":          operations.SigilAdd,
		"-":          operations.SigilSubstract,
		"*":          operations.SigilMultiply,
		"/":          operations.SigilDivide,
		"^":          operations.SigilPower,
	}
)

type Operation func(node *parser.Node, invoker func(node *parser.Node) (*parser.Value, error)) (*parser.Value, error)

func Invoke(node *parser.Node) (*parser.Value, error) {
	if node.Type == parser.TypeConstant {
		return node.Value, nil
	}

	if node.Type == parser.TypeOperation {
		operationName := *node.Value.StringVal

		operation, ok := operationsMap[operationName]
		if !ok {
			return nil, fmt.Errorf(`operand "%s" is not supported`, operationName)
		}

		return operation(node, Invoke)
	}

	return nil, fmt.Errorf(`onexpected operation %s`, node.Value)
}
