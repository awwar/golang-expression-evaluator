package operation

import (
	"fmt"

	"expression_parser/program"
	"expression_parser/utility"
)

func init() {
	AppendOperation(program.SKIP, Skip)
	AppendOperation(program.CSKIP, CSkip)
}

func Skip(pr *program.Program, stack *utility.Stack[program.Value], memo map[string]program.Value) error {
	op := pr.Current()

	n, ok := op.Params[0].(int)
	if !ok {
		return fmt.Errorf("SKIP param is not a int")
	}

	pr.Skip(n)

	return nil
}

func CSkip(pr *program.Program, stack *utility.Stack[program.Value], memo map[string]program.Value) error {
	operand, err := stack.Pop()
	if err != nil {
		return err
	}

	condition, err := operand.ToBoolean()
	if err != nil {
		return err
	}

	if !*condition {
		return nil
	}

	return Skip(pr, stack, memo)
}
