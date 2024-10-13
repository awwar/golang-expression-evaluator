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

func (p *Parser) Parse() (*Node, error) {
	var list []*Node

	for {
		token := p.stream.Get(p.currentPosition)

		if token == nil {
			return nil, fmt.Errorf("cant find token for position: %d", p.currentPosition)
		}

		if token.Type == tokenizer.TypeBrackets {
			endPosition := p.stream.SearchIdxOfClosedBracer(p.currentPosition)

			if endPosition == -1 {
				return nil, fmt.Errorf("cant find closed bracer for position: %d", p.currentPosition)
			}

			if p.currentPosition == endPosition-1 {
				return nil, fmt.Errorf("empty bracers detected at: %d", p.currentPosition)
			}

			subParser := New(p.stream, p.currentPosition+1, endPosition-1)

			subNode, err := subParser.Parse()

			if err != nil {
				return nil, err
			}

			subNode.Priority = 0

			list = append(list, subNode)

			p.currentPosition = endPosition
		}

		if token.Type == tokenizer.TypeOperation {
			node := &Node{Value: token, Priority: OperationPriority[token.Value]}

			list = append(list, node)
		}

		if token.Type == tokenizer.TypeNumber {
			list = append(list, &Node{Value: token, Priority: 0})
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

			if item.Priority != targetPriority {
				continue
			}

			if item.IsFilled() {
				continue
			}

			if i == 0 || i >= len(list) {
				return nil, fmt.Errorf("cant use infix operator without left or right part at: %d", p.currentPosition)
			}

			item.Left = list[i-1]
			item.Right = list[i+1]

			list = slices.Replace(list, i-1, i+2, item)

			i = i - 2
		}

		if targetPriority == 0 {
			break
		}

		targetPriority--
	}

	if len(list) != 1 {
		return nil, fmt.Errorf("after all transformations node list contains invalid amount of nodes: %d", len(list))
	}

	return list[0], nil
}
