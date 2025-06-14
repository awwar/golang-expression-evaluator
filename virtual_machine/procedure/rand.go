package procedure

import (
	"fmt"
	"math/rand"

	"expression_parser/parser"
	"expression_parser/utility"
)

func init() {
	AppendOperation("rand", Rand)
}

func Rand(argc int, stack *utility.Stack[parser.Value]) error {
	if argc != 0 {
		return fmt.Errorf("rand() not accepted any argument")
	}

	stack.Push(&parser.Value{Type: parser.Float, FloatVal: utility.AsPtr(rand.Float64())})

	return nil
}
