package parser

import (
	"fmt"

	"expression_parser/tokenizer"
	"expression_parser/utility"
)

var (
	ProcedureParserMap = map[string]ProcedureParser{}
)

func AddProcedureParser(name string, parser ProcedureParser) {
	ProcedureParserMap[name] = parser
}

type ProcedureParser interface {
	Parse(token *tokenizer.Token, pr *Parser) (*Node, error)
}

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

func (p *Parser) ParseProgram() (*Node, error) {
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

	return CreateAsProgram(list.Result()), nil
}

func (p *Parser) subparseExpressions() ([]*Node, error) {
	list := &NodeList{}

	for {
		token := p.stream.Get(p.currentPosition)
		if token == nil {
			lastToken := p.stream.Get(p.currentPosition - 1)

			return nil, p.error(lastToken.Position, "cant find token")
		}

		if token.Type == tokenizer.TypeSemicolon {
			subParser := New(p.stream, p.currentPosition+1, p.lastPosition)

			subNodes, err := subParser.subparseExpressions()
			if err != nil {
				return nil, err
			}

			list.Push(subNodes...)

			break
		} else if token.Type == tokenizer.TypeWord && token.StartsWith("#") {
			list.Push(CreateAsFlowLink(token.Value, token.Position))
		} else if token.Type == tokenizer.TypeWord && token.StartsWith("$") {
			list.Push(CreateAsVariable(token.Value, token.Position))
		} else if token.Type == tokenizer.TypeWord {
			subNodes, err := p.SubparseListInBracers(-1)
			if err != nil {
				return nil, err
			}

			list.Push(CreateAsOperation(token.Value, subNodes, token.Position))
		} else if token.Type == tokenizer.TypeBrackets {
			p.currentPosition--
			subNode, err := p.SubparseOneInBracers()
			if err != nil {
				return nil, err
			}

			subNode.SetPriority(0)

			list.Push(subNode)
		} else if token.Type == tokenizer.TypeOperation {
			list.Push(CreateAsOperation(token.Value, make([]*Node, 2), token.Position))
		} else if token.Type == tokenizer.TypeNumber {
			list.Push(CreateAsNumber(token.Value, token.Position))
		} else if token.Type == tokenizer.TypeString {
			list.Push(CreateAsString(token.Value, token.Position))
		} else {
			return nil, p.error(token.Position, "unknown token")
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
				return nil, p.error(currentNode.TokenPosition, err.Error())
			}

			if isReplaced {
				break
			}
		}
	}

	return list.Result(), nil
}

func (p *Parser) SubparseOneInBracers() (*Node, error) {
	subNodes, err := p.SubparseListInBracers(1)
	if err != nil {
		return nil, err
	}

	return subNodes[0], nil
}

func (p *Parser) SubparseVariableName() (*Node, error) {
	p.currentPosition++
	token := p.stream.Get(p.currentPosition)

	if token.Type != tokenizer.TypeWord {
		return nil, p.error(token.Position, "variable declaration must start with node name")
	}

	if !token.StartsWith("$") {
		return nil, p.error(token.Position, "variable declaration must start with $")
	}

	return CreateAsVariable(token.Value, token.Position), nil
}

func (p *Parser) SubparseFlowLink() (*Node, error) {
	p.currentPosition++
	token := p.stream.Get(p.currentPosition)

	if token.Type != tokenizer.TypeWord {
		return nil, p.error(token.Position, "hash links declaration must start with node name")
	}

	if !token.StartsWith("#") {
		return nil, p.error(token.Position, "hash links declaration must start with #")
	}

	return CreateAsFlowLink(token.Value, token.Position), nil
}

func (p *Parser) subparseFlowDeclaration() (*Node, error) {
	token := p.stream.Get(p.currentPosition)

	if token == nil {
		return nil, p.error(p.lastPosition, "cant find token")
	}

	if token.Type != tokenizer.TypeWord {
		return nil, p.error(token.Position, "hash links declaration must start with node name")
	}

	if token.StartsWith("#") {
		subNodes := NodeList{}

		isFlowDeclaration := false

		nextToken := p.stream.Get(p.currentPosition + 1)

		if nextToken.Type == tokenizer.TypeBrackets {
			args, err := p.SubparseListInBracers(-1)
			if err != nil {
				return nil, err
			}

			subNodes.Push(args...)

			returnParam, err := p.subparseWord()
			if err != nil {
				return nil, err
			}

			subNodes.Push(returnParam)

			isFlowDeclaration = true
		}

		for {
			nextToken = p.stream.Get(p.currentPosition + 1)

			if nextToken == nil || nextToken.StartsWith("#") {
				break
			}

			subNodeNode, err := p.subparseNode()
			if err != nil {
				return nil, err
			}

			subNodes.Push(subNodeNode)
		}

		if isFlowDeclaration {
			return CreateAsFlowDeclaration(token.Value, subNodes.Result(), token.Position), nil
		} else {
			return CreateAsFlowBranchesDeclaration(token.Value, subNodes.Result(), token.Position), nil
		}

	}

	return nil, p.error(token.Position, "flow declaration must start with # and has argument and return value")
}

func (p *Parser) subparseNode() (*Node, error) {
	p.currentPosition++
	token := p.stream.Get(p.currentPosition)

	if token.Type != tokenizer.TypeWord {
		return nil, p.error(token.Position, "node declaration must start with node name")
	}

	proc, ok := ProcedureParserMap[token.Value]
	if !ok {
		return nil, p.error(token.Position, fmt.Sprintf("unknown operation %s", token.Value))
	}

	return proc.Parse(token, p)
}

func (p *Parser) SubparseListInBracers(n int) ([]*Node, error) {
	p.currentPosition++

	openBracer := p.stream.Get(p.currentPosition)

	if openBracer == nil || openBracer.Type != tokenizer.TypeBrackets {
		errorToken := p.stream.Get(p.currentPosition - 1)

		return nil, p.error(errorToken.Position, "word token uses only in function context")
	}

	endPosition := p.stream.SearchIdxOfClosedBracer(p.currentPosition)

	if endPosition == -1 {
		currentToken := p.stream.Get(p.currentPosition)

		return nil, p.error(currentToken.Position, "cant find closed bracket")
	}

	subNodes := []*Node{}
	var err error

	if p.currentPosition != endPosition-1 {
		subParser := New(p.stream, p.currentPosition+1, endPosition-1)

		subNodes, err = subParser.subparseExpressions()
		if err != nil {
			return nil, err
		}
	}

	if n >= 0 && len(subNodes) != n {
		return nil, p.error(p.currentPosition+1, fmt.Sprintf("expected %d nodes, got %d", n, len(subNodes)))
	}

	p.currentPosition = endPosition

	return subNodes, nil
}

func (p *Parser) subparseWord() (*Node, error) {
	p.currentPosition++

	word := p.stream.Get(p.currentPosition)

	if word == nil || word.Type != tokenizer.TypeWord {
		return nil, p.error(p.lastPosition, "expected word token")
	}

	return CreateAsConstant(word.Value, p.lastPosition), nil
}

func (p *Parser) error(pos int, message string) error {
	return utility.NewError(pos, p.stream.Expression, message)
}
