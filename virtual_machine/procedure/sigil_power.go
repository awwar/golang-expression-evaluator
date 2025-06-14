package procedure

import (
	"errors"

	"expression_parser/parser"
	"expression_parser/utility"
)

func init() {
	AppendOperation("^", SigilPower)
}

func SigilPower(argc int, stack *utility.Stack[parser.Value]) error {
	if argc != 2 {
		return errors.New("sigil power: wrong number of arguments")
	}

	secondOperand, err := stack.Pop()
	if err != nil {
		return err
	}

	firstOperand, err := stack.Pop()
	if err != nil {
		return err
	}

	newV, err := firstOperand.Power(secondOperand)
	if err != nil {
		return err
	}

	stack.Push(newV)

	return nil
}
