package procedure

import (
	"errors"
	"strings"

	"expression_parser/program"
	"expression_parser/utility"
)

func init() {
	AppendProcedure("trim_space", &TrimSpace{})
}

type TrimSpace struct {
	CommonProcedure
}

func (t *TrimSpace) Execute(argc int, stack *utility.Stack[program.Value]) error {
	if argc != 1 {
		return errors.New("sigil add: wrong number of arguments")
	}

	firstOperand, err := stack.Pop()
	if err != nil {
		return err
	}

	strVal := strings.TrimSpace(*firstOperand.StringVal)

	stack.Push(program.NewString(strVal))

	return nil
}
