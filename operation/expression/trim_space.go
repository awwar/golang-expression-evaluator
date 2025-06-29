package expression

import (
	"errors"
	"strings"

	"expression_parser/operation"
	"expression_parser/program"
	"expression_parser/utility"
)

func init() {
	operation.AppendExpression("trim_space", &TrimSpace{})
}

type TrimSpace struct {
	operation.CommonExpression
}

func (t *TrimSpace) Execute(argc int, stack *utility.Stack[program.Value]) error {
	if argc != 1 {
		return errors.New("sigil add: wrong number of arguments")
	}

	firstOperand, err := stack.Pop()
	if err != nil {
		return err
	}

	strVal := strings.TrimSpace(firstOperand.String())

	stack.Push(program.NewString(strVal))

	return nil
}
