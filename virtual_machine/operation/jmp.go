package operation

import (
	"fmt"

	"expression_parser/program"
	"expression_parser/utility"
)

func init() {
	AppendOperation(program.JMP, Jmp)
}

func Jmp(pr *program.Program, _ *utility.Stack[program.Value], _ map[string]*program.Value) error {
	op := pr.Get()

	//pr.Skip(1)
	pr.TraceBack()

	markName, ok := op.Params[0].(string)
	if !ok {
		return fmt.Errorf("JMP markName param is not a string")
	}

	return pr.ToMark(markName)
}
