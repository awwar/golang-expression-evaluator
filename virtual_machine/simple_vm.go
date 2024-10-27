package virtual_machine

import (
	"expression_parser/parser"
	//"expression_parser/tokenizer"
	//"fmt"
	//"math"
	//"strconv"
)

func Invoke(node *parser.Node) (float64, error) {

	return 0.0, nil
	//if node.Value.Type == tokenizer.TypeNumber {
	//	value, err := strconv.ParseFloat(node.Value.Value, 32)
	//
	//	return value, err
	//}
	//
	//firstOperand, err := Invoke(node.Left)
	//
	//if err != nil {
	//	return 0.0, err
	//}
	//
	//secondOperand, err := Invoke(node.Right)
	//
	//if err != nil {
	//	return 0.0, err
	//}
	//
	//switch node.Value.Value {
	//case "+":
	//	return firstOperand + secondOperand, nil
	//case "-":
	//	return firstOperand - secondOperand, nil
	//case "*":
	//	return firstOperand * secondOperand, nil
	//case "/":
	//	return firstOperand / secondOperand, nil
	//case "^":
	//	return math.Pow(firstOperand, secondOperand), nil
	//default:
	//	return 0.0, fmt.Errorf(`operand type "%s" is not supported`, node.Value.Value)
	//}
}
