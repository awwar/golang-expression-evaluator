package parser

import (
	"expression_parser/tokenizer"
	"fmt"
	"strings"
)

type Node struct {
	Value    *tokenizer.Token
	Left     *Node
	Right    *Node
	Priority int
}

func (n *Node) IsFilled() bool {
	return n.Value.Type != tokenizer.TypeOperation || (n.Left != nil && n.Right != nil)
}

func (n *Node) String() string {
	return n.toStringWithIndention(0)
}

func (n *Node) toStringWithIndention(indent int) string {
	stringIndent := strings.Repeat("      ", indent)

	leftGraph := ""

	if n.Left != nil {
		leftGraph = fmt.Sprintf("%s└── L %s", stringIndent, n.Left.toStringWithIndention(indent+1))
	}

	rightGraph := ""

	if n.Right != nil {
		rightGraph = fmt.Sprintf("%s└── R %s", stringIndent, n.Right.toStringWithIndention(indent+1))
	}

	return fmt.Sprintf("%s\n%s%s", n.Value.Value, leftGraph, rightGraph)
}
