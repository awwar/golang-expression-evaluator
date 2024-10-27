package parser

import (
	"fmt"
	"strings"
)

type OperationNode struct {
	Operation string
	Left      Node
	Right     Node
	Priority  int
}

func (o *OperationNode) String(indent int) string {
	stringIndent := strings.Repeat("      ", indent)

	leftGraph := ""

	if o.Left != nil {
		leftGraph = fmt.Sprintf("%s└── L %s", stringIndent, o.Left.String(indent+1))
	}

	rightGraph := ""

	if o.Right != nil {
		rightGraph = fmt.Sprintf("%s└── R %s", stringIndent, o.Right.String(indent+1))
	}

	return fmt.Sprintf("%s\n%s\n%s", o.Operation, leftGraph, rightGraph)
}

func (o *OperationNode) SetPriority(priority int) {
	o.Priority = priority
}

func (o *OperationNode) GetPriority() int {
	return o.Priority
}

func (o *OperationNode) IsFilled() bool {
	return o.Left != nil && o.Right != nil
}

func (o *OperationNode) Fill(left, right Node) {
	o.Left = left
	o.Right = right
}
