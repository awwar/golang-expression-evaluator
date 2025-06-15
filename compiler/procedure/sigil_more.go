package procedure

import (
	"errors"

	"expression_parser/program"
	"expression_parser/utility"
)

func init() {
	AppendProcedure(">", &SigilMore{})
}

type SigilMore struct {
	CommonProcedure
}

func (s *SigilMore) Execute(argc int, stack *utility.Stack[program.Value]) error {
	if argc != 2 {
		return errors.New("sigil more: wrong number of arguments")
	}

	secondOperand, err := stack.Pop()
	if err != nil {
		return err
	}

	firstOperand, err := stack.Pop()
	if err != nil {
		return err
	}

	newV, err := firstOperand.More(secondOperand)
	if err != nil {
		return err
	}

	stack.Push(newV)

	return nil
}
