package expression

import (
	"errors"
	"strings"

	"expression_parser/operation"
	"expression_parser/program"
	"expression_parser/utility"
)

func init() {
	operation.AppendExpression("uppercase", &Uppercase{})
}

type Uppercase struct {
	operation.CommonExpression
}

func (u *Uppercase) Execute(argc int, stack *utility.Stack[program.Value]) error {
	if argc != 1 {
		return errors.New("sigil add: wrong number of arguments")
	}

	firstOperand, err := stack.Pop()
	if err != nil {
		return err
	}

	strVal := strings.ToUpper(firstOperand.String())

	stack.Push(program.NewString(strVal))

	return nil
}
