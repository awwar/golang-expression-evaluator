package procedures

import (
	"errors"
	"strings"

	"expression_parser/parser/expression"
	"expression_parser/utility"
)

func Uppercase(argc int, stack *utility.Stack[expression.Value]) (*expression.Value, error) {
	if argc != 1 {
		return nil, errors.New("sigil add: wrong number of arguments")
	}

	firstOperand, err := stack.Pop()
	if err != nil {
		return nil, err
	}

	strVal := strings.ToUpper(*firstOperand.StringVal)

	return &expression.Value{
		Type:      expression.String,
		StringVal: &strVal,
	}, nil
}
