package procedure

import (
	"errors"

	"expression_parser/program"
	"expression_parser/utility"
)

func init() {
	AppendProcedure("=", &SigilEq{})
}

type SigilEq struct {
	CommonProcedure
}

func (s *SigilEq) Execute(argc int, stack *utility.Stack[program.Value]) error {
	if argc != 2 {
		return errors.New("sigil eq: wrong number of arguments")
	}

	secondOperand, err := stack.Pop()
	if err != nil {
		return err
	}

	firstOperand, err := stack.Pop()
	if err != nil {
		return err
	}

	newV, err := firstOperand.Eq(secondOperand)
	if err != nil {
		return err
	}

	stack.Push(newV)

	return nil
}
