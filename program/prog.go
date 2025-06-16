package program

import (
	"fmt"
	"strconv"
	"strings"

	"expression_parser/utility"
)

type OperationName string

const (
	PUSH  OperationName = "PUSH"
	CALL  OperationName = "CALL"
	MARK  OperationName = "MARK"
	JMP   OperationName = "JMP"
	VAR   OperationName = "VAR"
	CSKIP OperationName = "CSKIP"
	SKIP  OperationName = "SKIP"
)

type Operation struct {
	Name   OperationName
	Params []any
}

func (o Operation) String() string {
	prs := ""
	for _, param := range o.Params {
		switch v := param.(type) {
		case Value:
			prs += v.String() + " "
		default:
			prs += fmt.Sprintf("%v ", param)
		}
	}

	return fmt.Sprintf("%s %s", o.Name, prs)
}

func NewProgram() *Program {
	return &Program{
		operations: []*Operation{},
		trace:      utility.NewStack[int](),
	}
}

type Program struct {
	operations []*Operation
	trace      *utility.Stack[int]
	opIdx      int
}

func (p *Program) NewMark(markName string) {
	p.operations = append(p.operations, &Operation{
		Name:   MARK,
		Params: []any{markName},
	})
}

func (p *Program) NewPush(value Value) {
	p.operations = append(p.operations, &Operation{
		Name:   PUSH,
		Params: []any{value},
	})
}

func (p *Program) NewVariable(name string) {
	p.operations = append(p.operations, &Operation{
		Name:   VAR,
		Params: []any{name},
	})
}

func (p *Program) NewJMP(markName string) {
	p.operations = append(p.operations, &Operation{
		Name:   JMP,
		Params: []any{markName},
	})
}

func (p *Program) NewCSKP(num int) {
	p.operations = append(p.operations, &Operation{
		Name:   CSKIP,
		Params: []any{num},
	})
}

func (p *Program) NewSKP(num int) {
	p.operations = append(p.operations, &Operation{
		Name:   SKIP,
		Params: []any{num},
	})
}

func (p *Program) NewCall(name string, argsC int) {
	p.operations = append(p.operations, &Operation{
		Name:   CALL,
		Params: []any{name, argsC},
	})
}

func (p *Program) Next() {
	p.opIdx++

	if !p.IsEnd() {
		return
	}

	p.FinishBlock()
}

func (p *Program) Current() *Operation {
	if p.IsEnd() {
		return nil
	}

	return p.operations[p.opIdx]
}

func (p *Program) TraceBack() {
	idx := p.opIdx + 1

	p.trace.Push(&idx)
}

func (p *Program) Skip(n int) {
	p.opIdx += n
}

func (p *Program) ToProgramBegin() error {
	if err := p.ToMark("#MAIN"); err != nil {
		return err
	}

	return nil
}

func (p *Program) ToMark(name string) error {
	for i, op := range p.operations {
		if op.Name == MARK {
			markName, ok := op.Params[0].(string)
			if !ok {
				return fmt.Errorf("markName param is not a string")
			}

			if markName == name {
				p.opIdx = i

				return nil
			}
		}
	}

	return fmt.Errorf("markName %s is not found", name)
}

func (p *Program) FinishBlock() {
	lastOpIdx, err := p.trace.Pop()
	if err != nil {
		p.opIdx = len(p.operations)

		return
	}
	p.opIdx = *lastOpIdx
}

func (p *Program) String() string {
	sb := strings.Builder{}

	for i, op := range p.operations {
		sb.WriteString(fmt.Sprintf("%d: %s\n", i, op.String()))
	}

	return sb.String()
}

func (p *Program) StringStatement() string {
	return fmt.Sprintf("op: %d trace: %s", p.opIdx, p.trace.ToString(func(n int) string { return strconv.Itoa(n) }))
}

func (p *Program) IsEnd() bool {
	return p.opIdx > len(p.operations)-1
}
