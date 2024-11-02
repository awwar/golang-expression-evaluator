package parser

import (
	"expression_parser/tokenizer"
	"fmt"
	"slices"
)

var OperationPriority = map[string]int{"+": 0, "-": 0, "*": 1, "/": 1, "^": 2}

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
			numberNode := CreateAsConstant(token.Value)

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

			item := list[i]

			if item.GetPriority() != targetPriority {
				continue
			}

			if item.IsFilled() {
				continue
			}

			if i == 0 || i >= len(list) {
				return nil, fmt.Errorf("cant use infix operator without left or right part at: %d", p.currentPosition)
			}

			item.SetSubNode(0, list[i-1])
			item.SetSubNode(1, list[i+1])

			list = slices.Replace(list, i-1, i+2, item)

			i = i - 2
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

	if p.currentPosition == endPosition-1 {
		return nil, nil
	}

	subParser := New(p.stream, p.currentPosition+1, endPosition-1)

	subNodes, err := subParser.Parse()

	p.currentPosition = endPosition

	if err != nil {
		return nil, err
	}

	return subNodes, nil
}
