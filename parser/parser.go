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

func (p *Parser) ParseProgram() ([]*Node, *Error) {
	list := &NodeList{}

	for {
		node, err := p.subparseFlowDeclaration()
		if err != nil {
			return nil, err
		}

		list.Push(node)

		if p.currentPosition == p.lastPosition {
			break
		}

		p.currentPosition++
	}

	return list.Result(), nil
}

func (p *Parser) subparseExpression() ([]*Node, *Error) {
	list := &NodeList{}

	for {
		token := p.stream.Get(p.currentPosition)

		if token == nil {
			lastToken := p.stream.Get(p.currentPosition - 1)

			return nil, NewError(lastToken.Position, "cant find token")
		}

		if token.Type == tokenizer.TypeSemicolon {
			subParser := New(p.stream, p.currentPosition+1, p.lastPosition)

			subNodes, err := subParser.subparseExpression()
			if err != nil {
				return nil, err
			}

			list.Push(subNodes...)

			break
		}

		if token.Type == tokenizer.TypeWord {
			if token.StartsWith("$") {
				node := CreateAsVariable(token.Value, token.Position)

				list.Push(node)
			} else {
				subNodes, err := p.subparseListInBracers()
				if err != nil {
					return nil, err
				}

				node := CreateAsOperation(token.Value, subNodes, token.Position)

				list.Push(node)
			}
		}

		if token.Type == tokenizer.TypeBrackets {
			p.currentPosition--
			subNode, err := p.subparseOneInBracers()
			if err != nil {
				return nil, err
			}

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

func (p *Parser) subparseListInBracers() ([]*Node, *Error) {
	p.currentPosition++

	openBracer := p.stream.Get(p.currentPosition)

	if openBracer == nil || openBracer.Type != tokenizer.TypeBrackets {
		errorToken := p.stream.Get(p.currentPosition - 1)

		return nil, NewError(errorToken.Position, "word token uses only in function context")
	}

	endPosition := p.stream.SearchIdxOfClosedBracer(p.currentPosition)

	if endPosition == -1 {
		currentToken := p.stream.Get(p.currentPosition)

		return nil, NewError(currentToken.Position, "cant find closed bracer")
	}

	var subNodes []*Node
	var err *Error

	if p.currentPosition != endPosition-1 {
		subParser := New(p.stream, p.currentPosition+1, endPosition-1)

		subNodes, err = subParser.subparseExpression()
	}

	p.currentPosition = endPosition

	if err != nil {
		return nil, err
	}

	return subNodes, nil
}

func (p *Parser) subparseOneInBracers() (*Node, *Error) {
	subNodes, err := p.subparseListInBracers()
	if err != nil {
		return nil, err
	}

	if len(subNodes) != 1 {
		for _, rt := range subNodes {
			fmt.Println(rt.String(0))
		}

		token := p.stream.Get(p.currentPosition)

		return nil, NewError(token.Position-1, "stand-alone brackets should frame exactly one node")
	}

	return subNodes[0], nil
}

func (p *Parser) subparseNode() (*Node, *Error) {
	token := p.stream.Get(p.currentPosition)

	if token.Type != tokenizer.TypeWord {
		return nil, NewError(token.Position, "node declaration must start with node name")
	}

	switch token.Value {
	case "VAR":
		variableNode, err := p.subparseVariableName()
		if err != nil {
			return nil, err
		}
		expression, err := p.subparseOneInBracers()
		if err != nil {
			return nil, err
		}
		return CreateAsOperation(token.Value, []*Node{variableNode, expression}, token.Position), nil
	case "IF":
		expression, err := p.subparseOneInBracers()
		if err != nil {
			return nil, err
		}
		trueHashLink, err := p.subparseFlowLink()
		if err != nil {
			return nil, err
		}
		falseHashLinks, err := p.subparseFlowLink()
		if err != nil {
			return nil, err
		}
		return CreateAsOperation(token.Value, []*Node{expression, trueHashLink, falseHashLinks}, token.Position), nil
	case "OUT":
		expression, err := p.subparseOneInBracers()
		if err != nil {
			return nil, err
		}
		return CreateAsOperation(token.Value, []*Node{expression}, token.Position), nil
	}

	return nil, NewError(token.Position, fmt.Sprintf("unknown operation %s", token.Value))
}

func (p *Parser) subparseVariableName() (*Node, *Error) {
	p.currentPosition++
	token := p.stream.Get(p.currentPosition)

	if token.Type != tokenizer.TypeWord {
		return nil, NewError(token.Position, "variable declaration must start with node name")
	}

	if !token.StartsWith("$") {
		return nil, NewError(token.Position, "variable declaration must start with $")
	}

	return CreateAsVariable(token.Value, token.Position), nil
}

func (p *Parser) subparseFlowDeclaration() (*Node, *Error) {
	token := p.stream.Get(p.currentPosition)

	if token == nil {
		lastToken := p.stream.Get(p.currentPosition - 1)

		return nil, NewError(lastToken.Position, "cant find token")
	}

	if token.Type != tokenizer.TypeWord {
		return nil, NewError(token.Position, "hash links declaration must start with node name")
	}

	if !token.StartsWith("#") {
		return nil, NewError(token.Position, "flow declaration must start with #")
	}

	subNodes := NodeList{}
	for {
		p.currentPosition++
		nextToken := p.stream.Get(p.currentPosition)

		if nextToken == nil || nextToken.StartsWith("#") {
			p.currentPosition--
			break
		}

		subNodeNode, err := p.subparseNode()
		if err != nil {
			return nil, err
		}

		subNodes.Push(subNodeNode)
	}

	return CreateAsFlowDeclaration(token.Value, subNodes.Result(), token.Position), nil
}

func (p *Parser) subparseFlowLink() (*Node, *Error) {
	p.currentPosition++
	token := p.stream.Get(p.currentPosition)

	if token.Type != tokenizer.TypeWord {
		return nil, NewError(token.Position, "hash links declaration must start with node name")
	}

	if !token.StartsWith("#") {
		return nil, NewError(token.Position, "hash links declaration must start with #")
	}

	return CreateAsFlowLink(token.Value, token.Position), nil
}
