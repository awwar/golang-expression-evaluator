package procedures

import (
	"errors"
	"strings"

	"expression_parser/parser"
	"expression_parser/utility"
)

func TrimSpace(argc int, stack *utility.Stack[parser.Value]) (*parser.Value, error) {
	if argc != 1 {
		return nil, errors.New("sigil add: wrong number of arguments")
	}

	firstOperand, err := stack.Pop()
	if err != nil {
		return nil, err
	}

	strVal := strings.TrimSpace(*firstOperand.StringVal)

	return &parser.Value{
		Type:      parser.String,
		StringVal: &strVal,
	}, nil
}
