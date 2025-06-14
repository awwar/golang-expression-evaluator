package procedure

import (
	"errors"
	"strings"

	"expression_parser/parser"
	"expression_parser/utility"
)

func init() {
	AppendOperation("trim_space", TrimSpace)
}

func TrimSpace(argc int, stack *utility.Stack[parser.Value]) error {
	if argc != 1 {
		return errors.New("sigil add: wrong number of arguments")
	}

	firstOperand, err := stack.Pop()
	if err != nil {
		return err
	}

	strVal := strings.TrimSpace(*firstOperand.StringVal)

	stack.Push(&parser.Value{Type: parser.String, StringVal: &strVal})

	return nil
}
