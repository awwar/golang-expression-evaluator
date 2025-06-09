package operations

import (
	"strings"

	"expression_parser/parser"
)

func TrimSpace(node *parser.Node, invoker func(node *parser.Node) (*parser.Value, error)) (*parser.Value, error) {
	paramResult, err := invoker(node.Params[0])
	if err != nil {
		return nil, err
	}

	strVal := strings.TrimSpace(*paramResult.StringVal)

	return &parser.Value{
		Type:      parser.String,
		StringVal: &strVal,
	}, nil
}
