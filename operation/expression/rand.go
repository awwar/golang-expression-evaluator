package expression

import (
	"fmt"
	"math/rand"

	"expression_parser/operation"
	"expression_parser/program"
	"expression_parser/utility"
)

func init() {
	operation.AppendExpression("rand", &Rand{})
}

type Rand struct {
	operation.CommonExpression
}

func (r *Rand) Execute(argc int, stack *utility.Stack[program.Value]) error {
	if argc != 0 {
		return fmt.Errorf("rand() not accepted any argument")
	}

	stack.Push(program.NewFloat(rand.Float64()))

	return nil
}
