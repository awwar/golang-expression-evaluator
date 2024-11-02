package virtual_machine

import (
	"expression_parser/parser"
	"fmt"
	"strings"
)

func Invoke(node *parser.Node) (*Value, error) {
	result := &Value{Type: Integer, Value: "0"}

	if node.Type == parser.TypeOperation {
		if node.Value == "sum" {
			for _, paramNode := range node.Params {
				paramResult, err := Invoke(paramNode)

				if err != nil {
					return nil, err
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

		switch node.Value {
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

	if node.Type == parser.TypeConstant {
		value := &Value{
			Type:  Integer,
			Value: node.Value,
		}

		if strings.Contains(node.Value, ".") {
			value.Type = Float
		}

		return value, nil
	}

	return nil, fmt.Errorf(`onexpected operation %s`, node.Value)
}
