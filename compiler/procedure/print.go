package procedure

import (
	"fmt"

	"expression_parser/program"
	"expression_parser/utility"
)

func init() {
	AppendProcedure("PRINT", &Print{})
}

type Print struct {
	CommonProcedure
}

func (o *Print) Execute(argc int, stack *utility.Stack[program.Value]) error {
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

	fmt.Println(*operandAsString.StringVal)

	return nil
}
