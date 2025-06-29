package node

import (
	"fmt"

	"expression_parser/operation"
	"expression_parser/parser"
	"expression_parser/program"
	"expression_parser/tokenizer"
	"expression_parser/utility"
)

func init() {
	operation.AppendProcedure("PRINT", &Print{})
}

type Print struct {
	operation.CommonProcedure
}

func (p *Print) Parse(token *tokenizer.Token, pr *parser.Parser) (*parser.Node, error) {
	expression, err := pr.SubparseOneInBracers()
	if err != nil {
		return nil, err
	}

	return parser.CreateAsOperation(token.Value, []*parser.Node{expression}, token.Position), nil
}

func (p *Print) Execute(argc int, stack *utility.Stack[program.Value]) error {
	if argc != 1 {
		return fmt.Errorf("PRINT() accepted only one argument")
	}

	firstOperand, err := stack.Pop()
	if err != nil {
		return err
	}

	operandAsString, err := firstOperand.ToString()
	if err != nil {
		return err
	}

	fmt.Println(operandAsString.String())

	return nil
}
