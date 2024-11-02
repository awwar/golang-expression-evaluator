package virtual_machine

import (
	"expression_parser/parser"
	"fmt"
	"math"
	"strconv"
)

func Invoke(node *parser.Node) (float64, error) {
	if node.Type == parser.TypeOperation {
		if node.Value == "sum" {
			result := 0.0

			for _, paramNode := range node.Params {
				paramResult, err := Invoke(paramNode)

				if err != nil {
					return 0.0, err
				}

				result += paramResult
			}

			return result, nil
		}

		firstOperand, err := Invoke(node.Params[0])

		if err != nil {
			return 0.0, err
		}

		secondOperand, err := Invoke(node.Params[1])

		if err != nil {
			return 0.0, err
		}

		switch node.Value {
		case "+":
			return firstOperand + secondOperand, nil
		case "-":
			return firstOperand - secondOperand, nil
		case "*":
			return firstOperand * secondOperand, nil
		case "/":
			return firstOperand / secondOperand, nil
		case "^":
			return math.Pow(firstOperand, secondOperand), nil
		default:
			return 0.0, fmt.Errorf(`operand type "%s" is not supported`, node.Value)
		}
	}

	if node.Type == parser.TypeConstant {
		value, err := strconv.ParseFloat(node.Value, 32)

		return value, err
	}

	return 0.0, fmt.Errorf(`onexpected operator`)
}
