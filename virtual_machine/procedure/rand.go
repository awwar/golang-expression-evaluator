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

func Rand(argc int, _ *utility.Stack[parser.Value]) (*parser.Value, error) {
	if argc != 0 {
		return nil, fmt.Errorf("rand() not accepted any argument")
	}

	return &parser.Value{
		Type:      parser.Float,
		StringVal: nil,
		FloatVal:  utility.AsPtr(rand.Float64()),
		IntVal:    nil,
	}, nil
}
