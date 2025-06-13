package procedures

import (
	"expression_parser/parser/expression"
	"expression_parser/utility"
)

func Sum(argc int, stack *utility.Stack[expression.Value]) (*expression.Value, error) {
	var result *expression.Value
	for i := 0; i < argc; i++ {
		operand, err := stack.Pop()
		if err != nil {
			return nil, err
		}

		if result == nil {
			result = operand

			continue
		}

		result, err = result.Add(operand)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
