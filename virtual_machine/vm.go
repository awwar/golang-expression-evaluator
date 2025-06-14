package virtual_machine

import (
	"fmt"

	"expression_parser/compiler"
	"expression_parser/parser"
	"expression_parser/utility"
	"expression_parser/virtual_machine/procedure"
)

func Execute(program compiler.Program) error {
	ops := program.Read()
	stack := utility.NewStack[parser.Value]()
	context := map[string]*parser.Value{}
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
		case compiler.MARK:
			opI = len(ops)
		case compiler.CALL:
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

			if err := proc(argc, stack); err != nil {
				return fmt.Errorf("procedure `%s` returns error: %v", procedureName, err)
			}
		case compiler.PUSH:
			v, ok := op.Params[0].(parser.Value)
			if !ok {
				return fmt.Errorf("PUSH param is not a *value")
			}
			if v.Type == parser.Atom {
				stack.Push(context[*v.StringVal])
			} else {
				stack.Push(&v)
			}
		case compiler.VAR:
			v, ok := op.Params[0].(parser.Value)
			if !ok {
				return fmt.Errorf("PUSH param is not a *value")
			}
			operand, err := stack.Pop()
			if err != nil {
				return err
			}
			context[*v.StringVal] = operand
		case compiler.IF:
			operand, err := stack.Pop()
			if err != nil {
				return err
			}
			if operand.Type != parser.Boolean {
				return fmt.Errorf("IF condition is not a boolean")
			}
			trace.Push(utility.AsPtr(opI))
			if operand.BoolVal {
				trueMark, ok := op.Params[0].(parser.Value)
				if !ok {
					return fmt.Errorf("PUSH param is not a *value")
				}

				opI = getBodyOfMark(ops, *trueMark.StringVal)
			} else {
				falseMark, ok := op.Params[1].(parser.Value)
				if !ok {
					return fmt.Errorf("PUSH param is not a *value")
				}

				opI = getBodyOfMark(ops, *falseMark.StringVal)
			}
		}
	}

	if !stack.IsEmpty() {
		return fmt.Errorf("after work stack is not empty")
	}

	return nil
}

func getBodyOfMark(ops []*compiler.Operations, mark string) int {
	for i, op := range ops {
		if op.Name == compiler.MARK {
			markName, ok := op.Params[0].(parser.Value)
			if !ok {
				return len(ops)
			}

			if *markName.StringVal == mark {
				return i + 1
			}
		}
	}

	return len(ops)
}
