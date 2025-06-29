package expression

import (
	"errors"

	"expression_parser/operation"
	"expression_parser/program"
	"expression_parser/utility"
)

func init() {
	operation.AppendExpression("bool", &TypeConversion{
		Op: func(a program.Value) (program.Value, error) { return a.ToBoolean() },
	})

	operation.AppendExpression("float", &TypeConversion{
		Op: func(a program.Value) (program.Value, error) { return a.ToFloat() },
	})

	operation.AppendExpression("string", &TypeConversion{
		Op: func(a program.Value) (program.Value, error) { return a.ToString() },
	})

	operation.AppendExpression("void", &TypeConversion{
		Op: func(a program.Value) (program.Value, error) { return program.NewInteger(0), nil },
	})
}

type TypeConversion struct {
	operation.CommonExpression
	Op func(a program.Value) (program.Value, error)
}

func (s *TypeConversion) Execute(argc int, stack *utility.Stack[program.Value]) error {
	if argc != 1 {
		return errors.New("wrong number of arguments")
	}

	firstOperand, err := stack.Pop()
	if err != nil {
		return err
	}

	newV, err := s.Op(firstOperand)
	if err != nil {
		return err
	}

	stack.Push(newV)

	return nil
}
