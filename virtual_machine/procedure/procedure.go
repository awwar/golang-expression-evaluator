package procedure

import (
	"expression_parser/parser"
	"expression_parser/utility"
)

var (
	ProceduresMap = map[string]Procedure{}
)

func AppendOperation(name string, op Procedure) {
	ProceduresMap[name] = op
}

type Procedure func(argc int, stack *utility.Stack[parser.Value]) error
