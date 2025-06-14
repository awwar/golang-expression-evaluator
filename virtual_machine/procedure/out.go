package procedure

import (
	"fmt"

	"expression_parser/parser"
	"expression_parser/utility"
)

func init() {
	AppendOperation("OUT", Out)
}

func Out(argc int, stack *utility.Stack[parser.Value]) error {
	if argc != 1 {
		return fmt.Errorf("OUT() accepted only one argument")
	}

	firstOperand, err := stack.Pop()
	if err != nil {
		return err
	}

	operandAsString, err := firstOperand.ToString()
	if err != nil {
		return err
	}

	fmt.Printf("%s", *operandAsString.StringVal)

	return nil
}
