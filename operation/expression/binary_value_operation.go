package expression

import (
	"errors"

	"expression_parser/operation"
	"expression_parser/program"
	"expression_parser/utility"
)

func init() {
	operation.AppendExpression("+", &BinaryValueOperation{
		Op: func(a, b *program.Value) (*program.Value, error) { return a.Add(b) },
	})

	operation.AppendExpression("-", &BinaryValueOperation{
		Op: func(a, b *program.Value) (*program.Value, error) { return a.Subtract(b) },
	})

	operation.AppendExpression("/", &BinaryValueOperation{
		Op: func(a, b *program.Value) (*program.Value, error) { return a.Divide(b) },
	})

	operation.AppendExpression("*", &BinaryValueOperation{
		Op: func(a, b *program.Value) (*program.Value, error) { return a.Multiply(b) },
	})

	operation.AppendExpression("^", &BinaryValueOperation{
		Op: func(a, b *program.Value) (*program.Value, error) { return a.Multiply(b) },
	})

	operation.AppendExpression("=", &BinaryValueOperation{
		Op: func(a, b *program.Value) (*program.Value, error) { return a.Eq(b) },
	})

	operation.AppendExpression("<", &BinaryValueOperation{
		Op: func(a, b *program.Value) (*program.Value, error) { return a.Less(b) },
	})

	operation.AppendExpression(">", &BinaryValueOperation{
		Op: func(a, b *program.Value) (*program.Value, error) { return a.More(b) },
	})
}

type BinaryValueOperation struct {
	operation.CommonExpression
	Op func(a, b *program.Value) (*program.Value, error)
}

func (s *BinaryValueOperation) Execute(argc int, stack *utility.Stack[program.Value]) error {
	if argc != 2 {
		return errors.New("sigil add: wrong number of arguments")
	}

	secondOperand, err := stack.Pop()
	if err != nil {
		return err
	}

	firstOperand, err := stack.Pop()
	if err != nil {
		return err
	}

	newV, err := s.Op(firstOperand, secondOperand)
	if err != nil {
		return err
	}

	stack.Push(newV)

	return nil
}
