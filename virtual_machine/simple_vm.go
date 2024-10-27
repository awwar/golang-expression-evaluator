package virtual_machine

import (
	"expression_parser/parser"
	"fmt"
	"math"
	"strconv"
)

func Invoke(node parser.Node) (float64, error) {
	switch n := node.(type) {
	case *parser.FunctionNode:
		switch n.Name {
		case "sum":
			result := 0.0

			for _, paramNode := range n.Params {
				paramResult, err := Invoke(paramNode)

				if err != nil {
					return 0.0, err
				}

				result += paramResult
			}

			return result, nil
		default:
			return 0.0, fmt.Errorf(`function "%s" is not supported`, n.Name)
		}
	case *parser.OperationNode:
		firstOperand, err := Invoke(n.Left)

		if err != nil {
			return 0.0, err
		}

		secondOperand, err := Invoke(n.Right)

		if err != nil {
			return 0.0, err
		}

		switch n.Operation {
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
			return 0.0, fmt.Errorf(`operand type "%s" is not supported`, n.Operation)
		}
	case *parser.ValueNode:
		value, err := strconv.ParseFloat(n.Value, 32)

		return value, err
	default:
		return 0.0, fmt.Errorf(`onexpected operator`)
	}
}
