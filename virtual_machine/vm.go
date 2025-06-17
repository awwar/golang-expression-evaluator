package virtual_machine

import (
	"fmt"

	"expression_parser/program"
	"expression_parser/utility"
	"expression_parser/virtual_machine/operation"
)

func Execute(pr *program.Program) error {
	stack := utility.NewStack[program.Value]()
	memo := map[string]*program.Value{}

	if err := pr.ToProgramBegin(); err != nil {
		return err
	}

	for {
		pr.Next()

		op := pr.Current()
		if op == nil {
			break
		}

		opExecute, ok := operation.Map[op.Name]
		if !ok {
			return fmt.Errorf("operation `%s` is not found", op.Name)
		}

		err := opExecute(pr, stack, memo)
		if err != nil {
			return err
		}

		// debug(op, stack, pr)
	}

	if !stack.IsEmpty() {
		return fmt.Errorf("after work stack is not empty")
	}

	return nil
}

func debug(op *program.Operation, stack *utility.Stack[program.Value], pr *program.Program) {
	fmt.Println(
		fmt.Sprintf(
			"> %s %s %s",
			op.String(),
			stack.ToString(func(v program.Value) string { return v.String() }),
			pr.StringStatement(),
		),
	)
}
