package procedure

import (
	"errors"

	"expression_parser/parser"
	"expression_parser/utility"
)

func init() {
	AppendOperation(">", SigilMore)
}

func SigilMore(argc int, stack *utility.Stack[parser.Value]) (*parser.Value, error) {
	if argc != 2 {
		return nil, errors.New("sigil more: wrong number of arguments")
	}

	secondOperand, err := stack.Pop()
	if err != nil {
		return nil, err
	}

	firstOperand, err := stack.Pop()
	if err != nil {
		return nil, err
	}

	return firstOperand.More(secondOperand)
}
