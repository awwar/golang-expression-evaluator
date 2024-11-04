package virtual_machine

import (
	"expression_parser/parser"
	"fmt"
)

func Invoke(node *parser.Node) (*parser.Value, error) {
	if node.Type == parser.TypeConstant {
		return node.Value, nil
	}

	if node.Type == parser.TypeOperation {
		var result *parser.Value

		operationName := *node.Value.StringVal

		if operationName == "sum" {
			for _, paramNode := range node.Params {
				paramResult, err := Invoke(paramNode)

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

		firstOperand, err := Invoke(node.Params[0])

		if err != nil {
			return nil, err
		}

		secondOperand, err := Invoke(node.Params[1])

		if err != nil {
			return nil, err
		}

		switch operationName {
		case "+":
			result, err = firstOperand.Add(secondOperand)
		case "-":
			result, err = firstOperand.Subtract(secondOperand)
		case "*":
			result, err = firstOperand.Multiply(secondOperand)
		case "/":
			result, err = firstOperand.Divide(secondOperand)
		case "^":
			result, err = firstOperand.Power(secondOperand)
		default:
			return nil, fmt.Errorf(`operand type "%s" is not supported`, node.Value)
		}

		if err != nil {
			return nil, err
		}

		return result, nil
	}

	return nil, fmt.Errorf(`onexpected operation %s`, node.Value)
}
