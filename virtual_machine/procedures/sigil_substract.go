package procedures

import (
	"errors"

	"expression_parser/parser/expression"
	"expression_parser/utility"
)

func SigilSubstract(argc int, stack *utility.Stack[expression.Value]) (*expression.Value, error) {
	if argc != 2 {
		return nil, errors.New("sigil add: wrong number of arguments")
	}

	secondOperand, err := stack.Pop()
	if err != nil {
		return nil, err
	}

	firstOperand, err := stack.Pop()
	if err != nil {
		return nil, err
	}

	return firstOperand.Subtract(secondOperand)
}
