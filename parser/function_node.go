package parser

import (
	"fmt"
	"strings"
)

type FunctionNode struct {
	Name     string
	Params   []Node
	Priority int
}

func (f *FunctionNode) String(indent int) string {
	stringIndent := strings.Repeat("      ", indent)

	branches := ""

	for i, n := range f.Params {
		branches = branches + fmt.Sprintf("%s└── #%d %s\n", stringIndent, i, n.String(indent+1))
	}

	return fmt.Sprintf("%s\n%s", f.Name, branches)
}

func (f *FunctionNode) IsFilled() bool {
	for _, n := range f.Params {
		if !n.IsFilled() {
			return false
		}
	}

	return true
}

func (f *FunctionNode) SetPriority(priority int) {
	f.Priority = priority
}

func (f *FunctionNode) GetPriority() int {
	return f.Priority
}

func (f *FunctionNode) Fill(left, right Node) {
}
