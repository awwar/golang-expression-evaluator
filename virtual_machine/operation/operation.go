package operation

import (
	"expression_parser/program"
	"expression_parser/utility"
)

type Operation func(pr *program.Program, stack *utility.Stack[program.Value], memo map[string]program.Value) error

var (
	Map = map[program.OperationName]Operation{}
)

func AppendOperation(name program.OperationName, op Operation) {
	Map[name] = op
}
