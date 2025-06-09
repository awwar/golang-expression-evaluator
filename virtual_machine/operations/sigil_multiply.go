package operations

import (
	"expression_parser/parser"
)

func SigilMultiply(node *parser.Node, invoker func(node *parser.Node) (*parser.Value, error)) (*parser.Value, error) {
	firstOperand, err := invoker(node.Params[0])
	if err != nil {
		return nil, err
	}

	secondOperand, err := invoker(node.Params[1])
	if err != nil {
		return nil, err
	}

	return firstOperand.Multiply(secondOperand)
}
