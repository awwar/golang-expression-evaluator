package virtual_machine

import (
	"fmt"

	"expression_parser/compiler"
	"expression_parser/parser"
	"expression_parser/utility"
	"expression_parser/virtual_machine/procedure"
)

func Execute(program compiler.Program) (*parser.Value, error) {
	ops := program.Read()
	stack := utility.NewStack[parser.Value]()
	context := map[string]*parser.Value{}

	opI := getIndexOfMark(ops, "#MAIN") + 1
	for {
		if opI > len(ops)-1 {
			break
		}
		op := ops[opI]
		opI++

		switch op.Name {
		case compiler.MARK:
			opI = len(ops)
		case compiler.CALL:
			procedureName, ok := op.Params[0].(string)
			if !ok {
				return nil, fmt.Errorf("CALL 1 param is not a string\n")
			}
			argc, ok := op.Params[1].(int)
			if !ok {
				return nil, fmt.Errorf("CALL 2 param is not a int\n")
			}
			proc, ok := procedure.ProceduresMap[procedureName]
			if !ok {
				return nil, fmt.Errorf("CALL procedure `%s` is not found\n", procedureName)
			}

			value, err := proc(argc, stack)
			if err != nil {
				return nil, fmt.Errorf("procedure `%s` returns error: %v\n", procedureName, err)
			}
			stack.Push(value)
		case compiler.PUSH:
			v, ok := op.Params[0].(parser.Value)
			if !ok {
				return nil, fmt.Errorf("PUSH param is not a *value\n")
			}
			if v.Type == parser.Atom {
				stack.Push(context[*v.StringVal])
			} else {
				stack.Push(&v)
			}
		case compiler.VAR:
			v, ok := op.Params[0].(parser.Value)
			if !ok {
				return nil, fmt.Errorf("PUSH param is not a *value\n")
			}
			operand, err := stack.Pop()
			if err != nil {
				return nil, err
			}
			context[*v.StringVal] = operand
		case compiler.IF:
			operand, err := stack.Pop()
			if err != nil {
				return nil, err
			}
			if operand.Type != parser.Boolean {
				return nil, fmt.Errorf("IF condition is not a boolean\n")
			}
			if operand.BoolVal {
				trueMark, ok := op.Params[0].(parser.Value)
				if !ok {
					return nil, fmt.Errorf("PUSH param is not a *value\n")
				}
				opI = getIndexOfMark(ops, *trueMark.StringVal) + 1
			} else {
				falseMark, ok := op.Params[1].(parser.Value)
				if !ok {
					return nil, fmt.Errorf("PUSH param is not a *value\n")
				}
				opI = getIndexOfMark(ops, *falseMark.StringVal) + 1
			}
		}
	}

	v, err := stack.Pop()
	if err != nil {
		return nil, fmt.Errorf("stack.Top() returns error: %v\n", err)
	}

	if !stack.IsEmpty() {
		return nil, fmt.Errorf("after work stack is not empty\n")
	}

	return v, nil
}

func getIndexOfMark(ops []*compiler.Operations, mark string) int {
	for i, op := range ops {
		if op.Name == compiler.MARK {
			markName, ok := op.Params[0].(parser.Value)
			if !ok {
				return -1
			}

			if *markName.StringVal == mark {
				return i
			}
		}
	}

	return -1
}
