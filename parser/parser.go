package parser

import (
	"fmt"

	"expression_parser/tokenizer"
)

type Parser struct {
	firstPosition   int
	lastPosition    int
	currentPosition int

	stream *tokenizer.TokenStream
}

var transformers = []Transformer{
	UnsignedMultiplication,
	ValueNegation,
	SimpleMath,
	FloatValue,
	FunctionCalling,
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
	list := &NodeList{}

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

			list.Push(subNodes...)

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

			list.Push(node)
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

			list.Push(subNode)
		}

		if token.Type == tokenizer.TypeOperation {
			node := CreateAsOperation(token.Value, make([]*Node, 2), token.Position)

			list.Push(node)
		}

		if token.Type == tokenizer.TypeNumber {
			node := CreateAsNumber(token.Value, token.Position)

			list.Push(node)
		}

		if token.Type == tokenizer.TypeString {
			node := CreateAsString(token.Value, token.Position)

			list.Push(node)
		}

		if p.currentPosition == p.lastPosition {
			break
		}

		p.currentPosition++
	}

	targetPriority := 4 + 1

	for {
		list.Next()

		if list.IsEnd() {
			list.Rewind()

			if targetPriority == 0 {
				break
			}

			targetPriority--
		}

		currentNode := list.Current()

		if currentNode.GetPriority() != targetPriority {
			continue
		}

		currentNode.Deprioritize()

		if list.Left() == nil {
			continue
		}

		for _, transformer := range transformers {
			isReplaced, err := transformer(list)
			if err != nil {
				return nil, err
			}

			if isReplaced {
				break
			}
		}
	}

	return list.Result(), nil
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
