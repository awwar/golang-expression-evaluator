package procedures

import (
	"expression_parser/parser"
	"expression_parser/utility"
)

func Sum(argc int, stack *utility.Stack[parser.Value]) (*parser.Value, error) {
	var result *parser.Value
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
