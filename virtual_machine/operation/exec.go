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
	AppendOperation(program.EXEC, Exec)
}

func Exec(pr *program.Program, stack *utility.Stack[program.Value], _ map[string]program.Value) error {
	op := pr.Current()

	procedureName, ok := op.Params[0].(string)
	if !ok {
		return fmt.Errorf("EXEC 1 param is not a string")
	}

	argc, ok := op.Params[1].(int)
	if !ok {
		return fmt.Errorf("EXEC 2 param is not a int")
	}

	proc, ok := ExternalMethodMap[procedureName]
	if !ok {
		return fmt.Errorf("EXEC procedure `%s` is not found", procedureName)
	}

	if err := proc.Execute(argc, stack); err != nil {
		return fmt.Errorf("procedure `%s` returns error: %v", procedureName, err)
	}

	return nil
}
