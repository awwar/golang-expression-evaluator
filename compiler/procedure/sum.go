package procedure

import (
	"expression_parser/program"
	"expression_parser/utility"
)

func init() {
	AppendProcedure("sum", &Sum{})
}

type Sum struct {
	CommonProcedure
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
