package operations

import (
	"strings"

	"expression_parser/parser"
)

func Uppercase(node *parser.Node, invoker func(node *parser.Node) (*parser.Value, error)) (*parser.Value, error) {
	paramResult, err := invoker(node.Params[0])
	if err != nil {
		return nil, err
	}

	strVal := strings.ToUpper(*paramResult.StringVal)

	return &parser.Value{
		Type:      parser.String,
		StringVal: &strVal,
	}, nil
}
