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

var OperationPriority = map[string]int{"+": 0, "-": 0, "*": 1, "/": 1, "^": 2, ".": 0}

type Node struct {
	Type          int
	Value         *Value
	Params        []*Node
	Priority      int
	TokenPosition int
}

func CreateAsOperation(operation string, params []*Node, tokenPosition int) *Node {
	return &Node{
		Type: TypeOperation,
		Value: &Value{
			Type:      Atom,
			StringVal: &operation,
		},
		Params:        params,
		Priority:      OperationPriority[operation],
		TokenPosition: tokenPosition,
	}
}

func CreateAsNumber(value string, tokenPosition int) *Node {
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
		Type:          TypeConstant,
		Value:         &valueObject,
		Params:        make([]*Node, 0),
		Priority:      0,
		TokenPosition: tokenPosition,
	}
}

func CreateAsString(value string, tokenPosition int) *Node {
	return &Node{
		Type: TypeConstant,
		Value: &Value{
			Type:      String,
			StringVal: &value,
		},
		Params:        make([]*Node, 0),
		Priority:      0,
		TokenPosition: tokenPosition,
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

func (f *Node) SetPriority(priority int) {
	f.Priority = priority
}

func (f *Node) GetPriority() int {
	return f.Priority
}

func (f *Node) SetSubNode(offset int, node *Node) {
	f.Params[offset] = node
}

func (f *Node) PushNodeToHead(node *Node) {
	if f.Params[0] == nil {
		f.Params = make([]*Node, 0)
	}

	f.Params = append([]*Node{node}, f.Params...)
}

func (f *Node) IsFilled() bool {
	for _, n := range f.Params {
		if n == nil || !n.IsFilled() {
			return false
		}
	}

	return true
}

func (f *Node) IsMathematicalOperation() bool {
	if f.Type != TypeOperation {
		return false
	}

	_, ok := OperationPriority[*f.Value.StringVal]

	return ok
}
