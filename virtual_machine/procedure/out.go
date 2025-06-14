package procedure

import (
	"fmt"

	"expression_parser/parser"
	"expression_parser/utility"
)

func init() {
	AppendOperation("OUT", Out)
}

func Out(argc int, stack *utility.Stack[parser.Value]) (*parser.Value, error) {
	if argc != 1 {
		return nil, fmt.Errorf("OUT() accepted only one argument")
	}

	firstOperand, err := stack.Pop()
	if err != nil {
		return nil, err
	}

	operandAsString, err := firstOperand.ToString()
	if err != nil {
		return nil, err
	}

	return &parser.Value{
		Type:      parser.String,
		StringVal: utility.AsPtr(fmt.Sprintf("%s", *operandAsString.StringVal)),
	}, nil
}
