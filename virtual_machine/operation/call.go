package operation

import (
	"fmt"

	"expression_parser/program"
	"expression_parser/utility"
)

type ExternalMethod interface {
	Execute(argc int, stack *utility.Stack[program.Value]) error
}

var ExternalMethodMap = map[string]ExternalMethod{}

func AddExternalMethod(name string, method ExternalMethod) {
	ExternalMethodMap[name] = method
}

func init() {
	AppendOperation(program.CALL, Call)
}

func Call(pr *program.Program, stack *utility.Stack[program.Value], memo map[string]program.Value) error {
	op := pr.Current()

	procedureName, ok := op.Params[0].(string)
	if !ok {
		return fmt.Errorf("CALL 1 param is not a string")
	}

	argc, ok := op.Params[1].(int)
	if !ok {
		return fmt.Errorf("CALL 2 param is not a int")
	}

	proc, ok := ExternalMethodMap[procedureName]
	if !ok {
		return fmt.Errorf("CALL procedure `%s` is not found", procedureName)
	}

	if err := proc.Execute(argc, stack); err != nil {
		return fmt.Errorf("procedure `%s` returns error: %v", procedureName, err)
	}

	return nil
}
