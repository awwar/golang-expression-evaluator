package compiler

import (
	"fmt"
	"strings"

	"expression_parser/parser"
)

type OperationName int

const (
	PUSH OperationName = iota
	CALL OperationName = iota
	MARK OperationName = iota
	IF   OperationName = iota
	VAR  OperationName = iota
)

type Operations struct {
	Name     OperationName
	Params   []any
	Describe func() string
}

func NewProgram() *Program {
	return &Program{
		operations: []*Operations{},
	}
}

type Program struct {
	operations []*Operations
}

func (p *Program) NewMark(markName string) {
	p.operations = append(p.operations, &Operations{
		Name:     MARK,
		Params:   []any{markName},
		Describe: func() string { return fmt.Sprintf("MARK %s", markName) },
	})
}

func (p *Program) NewPush(value parser.Value) {
	p.operations = append(p.operations, &Operations{
		Name:     PUSH,
		Params:   []any{value},
		Describe: func() string { return fmt.Sprintf("PUSH %s", value.GoString()) },
	})
}

func (p *Program) NewVariable(name parser.Value) {
	p.operations = append(p.operations, &Operations{
		Name:     VAR,
		Params:   []any{name},
		Describe: func() string { return fmt.Sprintf("VAR %s", name.GoString()) },
	})
}

func (p *Program) NewIf(ifTrueMarkName string, ifFalseMarkName string) {
	p.operations = append(p.operations, &Operations{
		Name:     IF,
		Params:   []any{ifTrueMarkName, ifFalseMarkName},
		Describe: func() string { return fmt.Sprintf("IF %s %s", ifTrueMarkName, ifFalseMarkName) },
	})
}

func (p *Program) NewCall(name string, argsC int) {
	p.operations = append(p.operations, &Operations{
		Name:     CALL,
		Params:   []any{name, argsC},
		Describe: func() string { return fmt.Sprintf("CALL %s %d", name, argsC) },
	})
}

func (p *Program) Read() []*Operations {
	return p.operations
}

func (p *Program) String() string {
	sb := strings.Builder{}

	for _, op := range p.operations {
		sb.WriteString(op.Describe() + "\n")
	}

	return sb.String()
}
