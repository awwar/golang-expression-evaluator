package parser

import (
	"expression_parser/tokenizer"
	"fmt"
	"slices"
)

var OperationPriority = map[string]int{"+": 0, "-": 0, "*": 1, "/": 1, "^": 2, ".": 0}

type Parser struct {
	firstPosition   int
	lastPosition    int
	currentPosition int

	stream *tokenizer.TokenStream
}

func New(stream *tokenizer.TokenStream, firstPosition, lastPosition int) *Parser {
	return &Parser{
		stream:          stream,
		firstPosition:   firstPosition,
		lastPosition:    lastPosition,
		currentPosition: firstPosition,
	}
}

func NewFromStream(stream *tokenizer.TokenStream) *Parser {
	return New(stream, 0, stream.Length()-1)
}

func (p *Parser) Parse() ([]*Node, error) {
	var list []*Node

	for {
		token := p.stream.Get(p.currentPosition)

		if token == nil {
			return nil, fmt.Errorf("cant find token for position: %d", p.currentPosition)
		}

		if token.Type == tokenizer.TypeWord {
			if false == p.stream.NextTokenIsBracer(p.currentPosition) {
				return nil, fmt.Errorf("word token uses only in function context: %d", p.currentPosition)
			}

			p.currentPosition++

			subNodes, err := p.subparseBracers()

			if err != nil {
				return nil, err
			}

			node := CreateAsOperation(token.Value, subNodes, 0)

			list = append(list, node)
		}

		if token.Type == tokenizer.TypeBrackets {
			subNodes, err := p.subparseBracers()

			if err != nil {
				return nil, err
			}

			if len(subNodes) != 1 {
				return nil, fmt.Errorf("stand-alone brackets should frame exactly one node: %d", p.currentPosition)
			}

			subNode := subNodes[0]

			subNode.SetPriority(0)

			list = append(list, subNode)
		}

		if token.Type == tokenizer.TypeOperation {
			operationNode := CreateAsOperation(token.Value, make([]*Node, 2), OperationPriority[token.Value])

			list = append(list, operationNode)
		}

		if token.Type == tokenizer.TypeNumber {
			numberNode := CreateAsNumber(token.Value)

			list = append(list, numberNode)
		}

		if token.Type == tokenizer.TypeString {
			numberNode := CreateAsString(token.Value)

			list = append(list, numberNode)
		}

		if p.currentPosition == p.lastPosition {
			break
		}

		p.currentPosition++
	}

	targetPriority := 2

	for {
		i := -1
		for {
			i++

			if len(list) < 2 || i >= len(list) {
				break
			}

			currentNode := list[i]

			if currentNode.GetPriority() != targetPriority {
				continue
			}

			if currentNode.IsFilled() {
				continue
			}

			if i >= len(list) {
				return nil, fmt.Errorf("cant use infix operator without rightNode part at: %d", p.currentPosition)
			}

			var leftNode *Node

			if i > 0 {
				leftNode = list[i-1]
			}

			rightNode := list[i+1]

			if leftNode == nil {
				newNode, err := p.createNegativeNode(currentNode, rightNode)

				if err != nil {
					return nil, err
				}

				list = slices.Replace(list, 0, 2, newNode)

				continue
			}

			if rightNode.Value.IsMinusOrPlus() && !rightNode.IsFilled() {
				rightRight := list[i+2]

				newNode, err := p.createNegativeNode(rightNode, rightRight)

				if err != nil {
					return nil, err
				}

				list = slices.Replace(list, i+1, i+3, newNode)
				i = i - 1

				continue
			}

			if *currentNode.Value.StringVal != "." {
				currentNode.SetSubNode(0, leftNode)
				currentNode.SetSubNode(1, rightNode)

				list = slices.Replace(list, i-1, i+2, currentNode)

				i = i - 2

				continue
			}

			if leftNode.Value.IsNumber() && rightNode.Value.IsNumber() {
				newNode := CreateAsNumber(fmt.Sprintf("%d.%d", *leftNode.Value.IntVal, *rightNode.Value.IntVal))

				list = slices.Replace(list, i-1, i+2, newNode)

				i = i - 1
			} else {
				currentNode.Value = rightNode.Value
				currentNode.PushNodeToHead(leftNode)

				list = slices.Replace(list, i-1, i+2, currentNode)

				i = i - 1
			}
		}

		if targetPriority == 0 {
			break
		}

		targetPriority--
	}

	return list, nil
}

func (p *Parser) subparseBracers() ([]*Node, error) {
	endPosition := p.stream.SearchIdxOfClosedBracer(p.currentPosition)

	if endPosition == -1 {
		return nil, fmt.Errorf("cant find closed bracer for position: %d", p.currentPosition)
	}

	var subNodes []*Node
	var err error

	if p.currentPosition != endPosition-1 {
		subParser := New(p.stream, p.currentPosition+1, endPosition-1)

		subNodes, err = subParser.Parse()
	}

	p.currentPosition = endPosition

	if err != nil {
		return nil, err
	}

	return subNodes, nil
}

func (p *Parser) createNegativeNode(operationNode *Node, operandNode *Node) (*Node, error) {
	operation := *operationNode.Value.StringVal

	if operation == "-" {
		if operandNode.Value.IsNumber() {
			var minusValue int64 = -1
			value := Value{
				Type:      Integer,
				StringVal: nil,
				FloatVal:  nil,
				IntVal:    &minusValue,
			}

			multipliedValue, err := operandNode.Value.Multiply(&value)

			if err != nil {
				return nil, err
			}

			stringVal, err := multipliedValue.ToString()

			if err != nil {
				return nil, err
			}

			numberNode := CreateAsNumber(*stringVal.StringVal)

			return numberNode, nil
		}

		numberNode := CreateAsNumber("-1")

		newOperationNode := CreateAsOperation("*", make([]*Node, 2), OperationPriority["*"])
		newOperationNode.SetSubNode(0, numberNode)
		newOperationNode.SetSubNode(1, operandNode)

		return newOperationNode, nil
	} else if operation == "+" {
		return operandNode, nil
	}

	return nil, fmt.Errorf("unable to negate node with operation: %s", operation)
}
