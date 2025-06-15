package procedure

import (
	"expression_parser/parser"
	"expression_parser/utility"
)

func init() {
	AppendOperation("sum", Sum)
}

func Sum(argc int, stack *utility.Stack[parser.Value]) error {
	var result *parser.Value

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
