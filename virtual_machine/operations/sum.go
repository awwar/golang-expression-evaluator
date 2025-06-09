package operations

import (
	"expression_parser/parser"
)

func Sum(node *parser.Node, invoker func(node *parser.Node) (*parser.Value, error)) (*parser.Value, error) {
	var result *parser.Value

	for _, paramNode := range node.Params {
		paramResult, err := invoker(paramNode)
		if err != nil {
			return nil, err
		}

		if result == nil {
			result = paramResult

			continue
		}

		result, err = result.Add(paramResult)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
