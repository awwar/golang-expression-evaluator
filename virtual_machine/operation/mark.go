package operation

import (
	"expression_parser/program"
	"expression_parser/utility"
)

func init() {
	AppendOperation(program.MARK, Mark)
}

func Mark(pr *program.Program, _ *utility.Stack[program.Value], _ map[string]program.Value) error {
	pr.FinishBlock()
	pr.Skip(-1)

	return nil
}
