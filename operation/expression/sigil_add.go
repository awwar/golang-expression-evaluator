package expression

import (
	"errors"

	"expression_parser/operation"
	"expression_parser/program"
	"expression_parser/utility"
)

func init() {
	operation.AppendExpression("+", &SigilAdd{})
}

type SigilAdd struct {
	operation.CommonExpression
}

func (s *SigilAdd) Execute(argc int, stack *utility.Stack[program.Value]) error {
	if argc != 2 {
		return errors.New("sigil add: wrong number of arguments")
	}

	secondOperand, err := stack.Pop()
	if err != nil {
		return err
	}

	firstOperand, err := stack.Pop()
	if err != nil {
		return err
	}

	newV, err := firstOperand.Add(secondOperand)
	if err != nil {
		return err
	}

	stack.Push(newV)

	return nil
}
