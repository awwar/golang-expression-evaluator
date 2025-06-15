package procedure

import (
	"errors"

	"expression_parser/parser"
	"expression_parser/utility"
)

func init() {
	AppendOperation("/", SigilDivide)
}

func SigilDivide(argc int, stack *utility.Stack[parser.Value]) error {
	if argc != 2 {
		return errors.New("sigil divide: wrong number of arguments")
	}

	secondOperand, err := stack.Pop()
	if err != nil {
		return err
	}

	firstOperand, err := stack.Pop()
	if err != nil {
		return err
	}

	newV, err := firstOperand.Divide(secondOperand)
	if err != nil {
		return err
	}

	stack.Push(newV)

	return nil
}
