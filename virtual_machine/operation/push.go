package operation

import (
	"fmt"

	"expression_parser/program"
	"expression_parser/utility"
)

func init() {
	AppendOperation(program.PUSH, Push)
}

func Push(pr *program.Program, stack *utility.Stack[program.Value], memo map[string]program.Value) error {
	op := pr.Current()

	v, ok := op.Params[0].(program.Value)
	if !ok {
		return fmt.Errorf("PUSH param is not a *value")
	}

	if program.IsVariable(v) {
		stack.Push(memo[v.String()])
	} else {
		stack.Push(v)
	}

	return nil
}
