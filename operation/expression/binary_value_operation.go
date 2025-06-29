package expression

import (
	"errors"

	"expression_parser/operation"
	"expression_parser/program"
	"expression_parser/utility"
)

func init() {
	operation.AppendExpression("+", &BinaryValueOperation{Op: program.Add})
	operation.AppendExpression("-", &BinaryValueOperation{Op: program.Subtract})
	operation.AppendExpression("/", &BinaryValueOperation{Op: program.Divide})
	operation.AppendExpression("*", &BinaryValueOperation{Op: program.Multiply})
	operation.AppendExpression("^", &BinaryValueOperation{Op: program.Power})
	operation.AppendExpression("=", &BinaryValueOperation{Op: program.Eq})
	operation.AppendExpression("<", &BinaryValueOperation{Op: program.Less})
	operation.AppendExpression(">", &BinaryValueOperation{Op: program.More})
}

type BinaryValueOperation struct {
	operation.CommonExpression
	Op func(a, b program.Value) (program.Value, error)
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
