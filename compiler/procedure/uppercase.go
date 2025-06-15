package procedure

import (
	"errors"
	"strings"

	"expression_parser/program"
	"expression_parser/utility"
)

func init() {
	AppendOperation("uppercase", &Uppercase{})
}

type Uppercase struct {
	CommonProcedure
}

func (u *Uppercase) Execute(argc int, stack *utility.Stack[program.Value]) error {
	if argc != 1 {
		return errors.New("sigil add: wrong number of arguments")
	}

	firstOperand, err := stack.Pop()
	if err != nil {
		return err
	}

	strVal := strings.ToUpper(*firstOperand.StringVal)

	stack.Push(program.NewString(strVal))

	return nil
}
