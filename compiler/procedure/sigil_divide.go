package procedure

import (
	"errors"

	"expression_parser/program"
	"expression_parser/utility"
)

func init() {
	AppendOperation("/", &SigilDivide{})
}

type SigilDivide struct {
	CommonProcedure
}

func (s *SigilDivide) Execute(argc int, stack *utility.Stack[program.Value]) error {
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
