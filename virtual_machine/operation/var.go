package operation

import (
	"fmt"

	"expression_parser/program"
	"expression_parser/utility"
)

func init() {
	AppendOperation(program.VAR, Var)
}

func Var(pr *program.Program, stack *utility.Stack[program.Value], memo map[string]*program.Value) error {
	op := pr.Current()

	varName, ok := op.Params[0].(string)
	if !ok {
		return fmt.Errorf("VAR name is not a string")
	}

	operand, err := stack.Pop()
	if err != nil {
		return err
	}

	if memo[varName] != nil {
		return fmt.Errorf("VAR unable to redefine %s", varName)
	}

	memo[varName] = operand

	return nil
}
