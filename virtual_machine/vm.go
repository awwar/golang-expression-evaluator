package virtual_machine

import (
	"fmt"

	"expression_parser/compiler/procedure"
	"expression_parser/program"
	"expression_parser/utility"
)

func Execute(pr program.Program) error {
	ops := pr.Read()
	stack := utility.NewStack[program.Value]()
	context := map[string]*program.Value{}
	trace := utility.NewStack[int]()

	opI := getBodyOfMark(ops, "#MAIN")
	for {
		if opI > len(ops)-1 {
			lastOp, err := trace.Pop()
			if err != nil {
				break
			}
			opI = *lastOp
		}
		op := ops[opI]
		opI++

		switch op.Name {
		case program.MARK:
			opI = len(ops)
		case program.CALL:
			procedureName, ok := op.Params[0].(string)
			if !ok {
				return fmt.Errorf("CALL 1 param is not a string")
			}
			argc, ok := op.Params[1].(int)
			if !ok {
				return fmt.Errorf("CALL 2 param is not a int")
			}
			proc, ok := procedure.ProceduresMap[procedureName]
			if !ok {
				return fmt.Errorf("CALL procedure `%s` is not found", procedureName)
			}

			if err := proc.Execute(argc, stack); err != nil {
				return fmt.Errorf("procedure `%s` returns error: %v", procedureName, err)
			}
		case program.PUSH:
			v, ok := op.Params[0].(program.Value)
			if !ok {
				return fmt.Errorf("PUSH param is not a *value")
			}
			if v.IsAtom() {
				stack.Push(context[*v.StringVal])
			} else {
				stack.Push(&v)
			}
		case program.VAR:
			v, ok := op.Params[0].(program.Value)
			if !ok {
				return fmt.Errorf("PUSH param is not a *value")
			}
			operand, err := stack.Pop()
			if err != nil {
				return err
			}
			context[*v.StringVal] = operand
		case program.IF:
			trace.Push(utility.AsPtr(opI))

			operand, err := stack.Pop()
			if err != nil {
				return err
			}

			condition, err := operand.ToBoolean()
			if err != nil {
				return err
			}

			if *condition.BoolVal {
				trueMark, ok := op.Params[0].(string)
				if !ok {
					return fmt.Errorf("IF trueMark param is not a string")
				}

				opI = getBodyOfMark(ops, trueMark)
			} else {
				falseMark, ok := op.Params[1].(string)
				if !ok {
					return fmt.Errorf("IF falseMark param is not a string")
				}

				opI = getBodyOfMark(ops, falseMark)
			}
		}
	}

	if !stack.IsEmpty() {
		return fmt.Errorf("after work stack is not empty")
	}

	return nil
}

func getBodyOfMark(ops []*program.Operations, mark string) int {
	for i, op := range ops {
		if op.Name == program.MARK {
			markName, ok := op.Params[0].(string)
			if !ok {
				return len(ops)
			}

			if markName == mark {
				return i + 1
			}
		}
	}

	return len(ops)
}
