package parser

import (
	"expression_parser/tokenizer"
	"expression_parser/utility"
	"fmt"
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

func (p *Parser) Parse() (*Node, error) {
	stack := &utility.Stack[Node]{}

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

			if stack.IsEmpty() {
				stack.Push(subNode)
			} else {
				node, _ := stack.Top()

				if node.Value.Type == tokenizer.TypeNumber {
					return nil, fmt.Errorf("find another number on top of stack: %d", p.currentPosition)
				}

				if node.Left == nil {
					node.Left = subNode
				} else if node.Right == nil {
					node.Right = subNode
				} else {
					return nil, fmt.Errorf("find orphan bracers: %d", p.currentPosition)
				}
			}

			p.currentPosition = endPosition
		}

		if token.Type == tokenizer.TypeOperation {
			if stack.IsEmpty() {
				return nil, fmt.Errorf("cant use infix operator without left part at position: %d", p.currentPosition)
			}

			node, _ := stack.Top()

			if node.Value.Type == tokenizer.TypeNumber {
				stack.Pop()

				stack.Push(&Node{Value: token, Left: node})
			} else if node.Right == nil {
				node.Right = &Node{Value: token}
			} else {
				stack.Pop()

				stack.Push(&Node{Value: token, Left: node})
			}
		}

		if token.Type == tokenizer.TypeNumber {
			if stack.IsEmpty() {
				stack.Push(&Node{Value: token})
			} else {
				node, _ := stack.Top()

				if node.Value.Type == tokenizer.TypeNumber {
					return nil, fmt.Errorf("find another number on top of stack: %d", p.currentPosition)
				}

				if node.Left == nil {
					node.Left = &Node{Value: token}
				} else if node.Right == nil {
					node.Right = &Node{Value: token}
				} else {
					return nil, fmt.Errorf("find orphan number: %d", p.currentPosition)
				}
			}
		}

		if p.currentPosition == p.lastPosition {
			tree, _ := stack.Top()

			return tree, nil
		}

		p.currentPosition++
	}
}
