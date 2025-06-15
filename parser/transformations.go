package parser

import (
	"fmt"

	"expression_parser/program"
)

type Transformer func(list TransformableNodeList) (bool, error)

type TransformableNodeList interface {
	Current() *Node
	Left() *Node
	Right() *Node
	RightRight() *Node
	LeftLeft() *Node
	Replace(toLeft, toRight int, node *Node)
}

// UnsignedMultiplication 2 sin(20) -> (2 * sin(20))
func UnsignedMultiplication(list TransformableNodeList) (bool, error) {
	leftNode := list.Left()
	currentNode := list.Current()

	if leftNode == nil {
		return false, nil
	}

	if !currentNode.IsFunction() || !leftNode.IsNumber() {
		return false, nil
	}

	newNode := CreateAsOperation("*", make([]*Node, 2), currentNode.TokenPosition)
	newNode.SetSubNode(0, currentNode)
	newNode.SetSubNode(1, leftNode)

	newNode.Deprioritize()

	list.Replace(1, 1, newNode)

	return true, nil
}

// ValueNegation 1 + - 1 -> 1 + -1
func ValueNegation(list TransformableNodeList) (bool, error) {
	leftLeftNode := list.LeftLeft()
	leftNode := list.Left()
	currentNode := list.Current()

	if leftNode == nil {
		return false, nil
	}

	var isAllowedLeftBound = leftLeftNode == nil || leftLeftNode.IsMathematicalOperation()

	if !(currentNode.IsNegatable() && leftNode.IsMinusOrPlus() && isAllowedLeftBound) {
		return false, nil
	}

	newNode, err := createNegativeNode(leftNode, currentNode)

	if err != nil {
		return false, err
	}

	newNode.Deprioritize()

	list.Replace(1, 1, newNode)

	return true, nil
}

func SimpleMath(list TransformableNodeList) (bool, error) {
	leftNode := list.Left()
	currentNode := list.Current()
	rightNode := list.Right()

	if leftNode == nil || rightNode == nil {
		return false, nil
	}

	if !currentNode.IsMathematicalOperation() {
		return false, nil
	}

	currentNode.SetSubNode(0, leftNode)
	currentNode.SetSubNode(1, rightNode)

	list.Replace(1, 2, currentNode)

	return true, nil
}

func FloatValue(list TransformableNodeList) (bool, error) {
	leftNode := list.Left()
	currentNode := list.Current()
	rightNode := list.Right()

	if leftNode == nil || rightNode == nil {
		return false, nil
	}

	if !currentNode.IsCallOperation() {
		return false, nil
	}

	if !leftNode.IsNumber() || !rightNode.IsNumber() {
		return false, nil
	}

	strFloatNumber := fmt.Sprintf("%d.%d", *leftNode.Value.IntVal, *rightNode.Value.IntVal)
	newNode := CreateAsNumber(strFloatNumber, rightNode.TokenPosition)

	list.Replace(1, 2, newNode)

	newNode.Deprioritize()

	return true, nil
}

func FunctionCalling(list TransformableNodeList) (bool, error) {
	leftNode := list.Left()
	currentNode := list.Current()
	rightNode := list.Right()

	if leftNode == nil || rightNode == nil {
		return false, nil
	}

	if !currentNode.IsCallOperation() {
		return false, nil
	}

	currentNode.Value = rightNode.Value
	currentNode.SetOnlyChild(leftNode)

	list.Replace(1, 2, currentNode)

	return true, nil
}

func createNegativeNode(operationNode *Node, operandNode *Node) (*Node, error) {
	operation := *operationNode.Value.StringVal

	if operation == "+" {
		return operandNode, nil
	}

	if operation != "-" {
		return nil, fmt.Errorf("unable to negate node with operation %s", operation)
	}

	if operandNode.IsNumber() {
		var minusValue int64 = -1
		value := program.Value{
			ValueType: program.Integer,
			StringVal: nil,
			FloatVal:  nil,
			IntVal:    &minusValue,
		}

		multipliedValue, err := operandNode.Value.Multiply(&value)

		if err != nil {
			return nil, fmt.Errorf("negation value error: %s", err)
		}

		stringVal, err := multipliedValue.ToString()

		if err != nil {
			return nil, fmt.Errorf("negation value error: %s", err)
		}

		numberNode := CreateAsNumber(*stringVal.StringVal, operandNode.TokenPosition)

		return numberNode, nil
	}

	numberNode := CreateAsNumber("-1", operandNode.TokenPosition)

	newOperationNode := CreateAsOperation("*", make([]*Node, 2), operandNode.TokenPosition)
	newOperationNode.SetSubNode(0, numberNode)
	newOperationNode.SetSubNode(1, operandNode)

	return newOperationNode, nil
}
