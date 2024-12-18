package parser

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	TypeOperation = iota
	TypeConstant  = iota
)

type Node struct {
	Type     int
	Value    *Value
	Params   []*Node
	Priority int
}

func CreateAsOperation(operation string, params []*Node, priority int) *Node {
	return &Node{
		Type: TypeOperation,
		Value: &Value{
			Type:      Atom,
			StringVal: &operation,
		},
		Params:   params,
		Priority: priority,
	}
}

func CreateAsNumber(value string) *Node {
	valueObject := Value{}

	if strings.Contains(value, ".") {
		val, _ := strconv.ParseFloat(value, 64)

		valueObject.Type = Float
		valueObject.FloatVal = &val
	} else {
		val, _ := strconv.ParseInt(value, 0, 64)

		valueObject.Type = Integer
		valueObject.IntVal = &val
	}

	return &Node{
		Type:     TypeConstant,
		Value:    &valueObject,
		Params:   make([]*Node, 0),
		Priority: 0,
	}
}

func CreateAsString(value string) *Node {
	return &Node{
		Type: TypeConstant,
		Value: &Value{
			Type:      String,
			StringVal: &value,
		},
		Params:   make([]*Node, 0),
		Priority: 0,
	}
}

func (f *Node) String(indent int) string {
	stringIndent := strings.Repeat("      ", indent)

	branches := ""

	for i, n := range f.Params {
		branches = branches + fmt.Sprintf("%s└── #%d %s", stringIndent, i, n.String(indent+1))
	}

	return fmt.Sprintf("%s\n%s", f.Value, branches)
}

func (f *Node) IsFilled() bool {
	for _, n := range f.Params {
		if n == nil || !n.IsFilled() {
			return false
		}
	}

	return true
}

func (f *Node) SetPriority(priority int) {
	f.Priority = priority
}

func (f *Node) GetPriority() int {
	return f.Priority
}

func (f *Node) SetSubNode(offset int, node *Node) {
	f.Params[offset] = node
}
