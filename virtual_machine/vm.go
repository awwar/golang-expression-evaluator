package virtual_machine

import (
	"fmt"

	"expression_parser/compiler"
	"expression_parser/parser"
	"expression_parser/utility"
	"expression_parser/virtual_machine/procedures"
)

var (
	proceduresMap = map[string]Operation{
		"sum":        procedures.Sum,
		"uppercase":  procedures.Uppercase,
		"trim_space": procedures.TrimSpace,
		"+":          procedures.SigilAdd,
		"-":          procedures.SigilSubstract,
		"*":          procedures.SigilMultiply,
		"/":          procedures.SigilDivide,
		"^":          procedures.SigilPower,
	}
)

type Operation func(argc int, stack *utility.Stack[parser.Value]) (*parser.Value, error)

func Execute(program compiler.Program) (*parser.Value, error) {
	ops := program.Read()
	stack := utility.NewStack[parser.Value]()

	opI := 0
	for {
		if opI > len(ops)-1 {
			break
		}
		op := ops[opI]
		opI++

		switch op.Name {
		case compiler.CALL:
			procedureName, ok := op.Params[0].(string)
			if !ok {
				return nil, fmt.Errorf("CALL 1 param is not a string\n")
			}
			argc, ok := op.Params[1].(int)
			if !ok {
				return nil, fmt.Errorf("CALL 2 param is not a int\n")
			}
			procedure, ok := proceduresMap[procedureName]
			if !ok {
				return nil, fmt.Errorf("CALL procedure `%s` is not found\n", procedureName)
			}

			value, err := procedure(argc, stack)
			if err != nil {
				return nil, fmt.Errorf("procedure `%s` returns error: %v\n", procedureName, err)
			}
			stack.Push(value)
		case compiler.PUSH:
			v, ok := op.Params[0].(parser.Value)
			if !ok {
				return nil, fmt.Errorf("PUSH param is not a *value\n")
			}
			stack.Push(&v)
		}
	}

	v, err := stack.Pop()
	if err != nil {
		return nil, fmt.Errorf("stack.Top() returns error: %v\n", err)
	}

	return v, nil
}
