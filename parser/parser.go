package parser

import (
	"expression_parser/tokenizer"
	"fmt"
	"slices"
)

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

func (p *Parser) Parse() ([]*Node, *Error) {
	var list []*Node

	for {
		token := p.stream.Get(p.currentPosition)

		if token == nil {
			lastToken := p.stream.Get(p.currentPosition - 1)

			return nil, NewError(lastToken.Position, "cant find token")
		}

		if token.Type == tokenizer.TypeSemicolon {
			subParser := New(p.stream, p.currentPosition+1, p.lastPosition)

			subNodes, err := subParser.Parse()

			if err != nil {
				return nil, err
			}

			list = append(list, subNodes...)

			break
		}

		if token.Type == tokenizer.TypeWord {
			if false == p.stream.NextTokenIsBracer(p.currentPosition) {
				return nil, NewError(token.Position, "word token uses only in function context")
			}

			p.currentPosition++

			subNodes, err := p.subparseBracers()

			if err != nil {
				return nil, err
			}

			node := CreateAsOperation(token.Value, subNodes, token.Position)

			list = append(list, node)
		}

		if token.Type == tokenizer.TypeBrackets {
			subNodes, err := p.subparseBracers()

			if err != nil {
				return nil, err
			}

			if len(subNodes) != 1 {
				for _, rt := range subNodes {
					fmt.Println(rt.String(0))
				}

				return nil, NewError(token.Position-1, "stand-alone brackets should frame exactly one node")
			}

			subNode := subNodes[0]

			subNode.SetPriority(0)

			list = append(list, subNode)
		}

		if token.Type == tokenizer.TypeOperation {
			operationNode := CreateAsOperation(token.Value, make([]*Node, 2), token.Position)

			list = append(list, operationNode)
		}

		if token.Type == tokenizer.TypeNumber {
			numberNode := CreateAsNumber(token.Value, token.Position)

			list = append(list, numberNode)
		}

		if token.Type == tokenizer.TypeString {
			numberNode := CreateAsString(token.Value, token.Position)

			list = append(list, numberNode)
		}

		if p.currentPosition == p.lastPosition {
			break
		}

		p.currentPosition++
	}

	targetPriority := 4 + 1

	for {
		i := -1

		if targetPriority == 0 {
			break
		}

		targetPriority--

		for {
			i++

			if len(list) < 2 || i >= len(list) {
				break
			}

			currentNode := list[i]

			if currentNode.GetPriority() != targetPriority {
				continue
			}

			currentNode.Deprioritize()

			var leftNode *Node
			var rightNode *Node

			if i > 0 {
				leftNode = list[i-1]
			}

			if i < len(list)-1 {
				rightNode = list[i+1]
			}

			if leftNode == nil {
				continue
			}

			// 2 sin(20) ~ (2 * sin(20))
			if currentNode.IsFunction() && leftNode.IsNumber() {
				newNode := CreateAsOperation("*", make([]*Node, 2), currentNode.TokenPosition)
				newNode.SetSubNode(0, currentNode)
				newNode.SetSubNode(1, leftNode)

				newNode.Deprioritize()

				list = slices.Replace(list, i-1, i+1, newNode)

				i = i - 1

				continue
			}

			// 1 + - 1 ~ 1 + -1
			if currentNode.IsNegatable() && leftNode.IsMinusOrPlus() && (i < 2 || list[i-2].IsMathematicalOperation()) {
				newNode, err := p.createNegativeNode(leftNode, currentNode)

				if err != nil {
					return nil, err
				}

				newNode.Deprioritize()

				list = slices.Replace(list, i-1, i+1, newNode)

				i = i - 1

				continue
			}

			if rightNode == nil {
				continue
			}

			if currentNode.IsMathematicalOperation() {
				currentNode.SetSubNode(0, leftNode)
				currentNode.SetSubNode(1, rightNode)

				list = slices.Replace(list, i-1, i+2, currentNode)

				i = i - 2

				continue
			}

			if currentNode.IsCallOperation() {
				if leftNode.IsNumber() && rightNode.IsNumber() {
					strFloatNumber := fmt.Sprintf("%d.%d", *leftNode.Value.IntVal, *rightNode.Value.IntVal)
					newNode := CreateAsNumber(strFloatNumber, rightNode.TokenPosition)

					list = slices.Replace(list, i-1, i+2, newNode)

					i = i - 1

					newNode.Deprioritize()

					continue
				}

				currentNode.Value = rightNode.Value
				currentNode.PushNodeToHead(leftNode)

				list = slices.Replace(list, i-1, i+2, currentNode)

				i = i - 1

				continue
			}
		}
	}

	return list, nil
}

func (p *Parser) subparseBracers() ([]*Node, *Error) {
	endPosition := p.stream.SearchIdxOfClosedBracer(p.currentPosition)

	if endPosition == -1 {
		currentToken := p.stream.Get(p.currentPosition)

		return nil, NewError(currentToken.Position, "cant find closed bracer")
	}

	var subNodes []*Node
	var err *Error

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

func (p *Parser) createNegativeNode(operationNode *Node, operandNode *Node) (*Node, *Error) {
	operation := *operationNode.Value.StringVal

	if operation == "-" {
		if operandNode.IsNumber() {
			var minusValue int64 = -1
			value := Value{
				Type:      Integer,
				StringVal: nil,
				FloatVal:  nil,
				IntVal:    &minusValue,
			}

			multipliedValue, err := operandNode.Value.Multiply(&value)

			if err != nil {
				return nil, NewError(operationNode.TokenPosition, "negation value error: %s", err)
			}

			stringVal, err := multipliedValue.ToString()

			if err != nil {
				return nil, NewError(operationNode.TokenPosition, "negation value error: %s", err)
			}

			numberNode := CreateAsNumber(*stringVal.StringVal, operandNode.TokenPosition)

			return numberNode, nil
		}

		numberNode := CreateAsNumber("-1", operandNode.TokenPosition)

		newOperationNode := CreateAsOperation("*", make([]*Node, 2), operandNode.TokenPosition)
		newOperationNode.SetSubNode(0, numberNode)
		newOperationNode.SetSubNode(1, operandNode)

		return newOperationNode, nil
	} else if operation == "+" {
		return operandNode, nil
	}

	return nil, NewError(operationNode.TokenPosition, "unable to negate node with operation %s", operation)
}
