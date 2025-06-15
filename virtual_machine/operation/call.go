package operation

import (
	"fmt"

	"expression_parser/compiler/procedure"
	"expression_parser/program"
	"expression_parser/utility"
)

func init() {
	AppendOperation(program.CALL, Call)
}

func Call(pr *program.Program, stack *utility.Stack[program.Value], memo map[string]*program.Value) error {
	op := pr.Get()

	procedureName, ok := op.Params[0].(string)
	if !ok {
		return fmt.Errorf("CALL 1 param is not a string")
	}
	argc, ok := op.Params[1].(int)
	if !ok {
		return fmt.Errorf("CALL 2 param is not a int")
	}
	proc, ok := procedure.Map[procedureName]
	if !ok {
		return fmt.Errorf("CALL procedure `%s` is not found", procedureName)
	}

	if err := proc.Execute(argc, stack); err != nil {
		return fmt.Errorf("procedure `%s` returns error: %v", procedureName, err)
	}

	return nil
}
