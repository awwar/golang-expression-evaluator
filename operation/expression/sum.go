package expression

import (
	"expression_parser/operation"
	"expression_parser/program"
	"expression_parser/utility"
)

func init() {
	operation.AppendExpression("sum", &Sum{})
}

type Sum struct {
	operation.CommonExpression
}

func (s *Sum) Execute(argc int, stack *utility.Stack[program.Value]) error {
	var result *program.Value

	for i := 0; i < argc; i++ {
		operand, err := stack.Pop()
		if err != nil {
			return err
		}
		if result == nil {
			result = operand

			continue
		}

		result, err = operand.Add(result)
		if err != nil {
			return err
		}
	}

	stack.Push(result)

	return nil
}
