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

var OperationPriority = map[string]int{"+": 1, "-": 2, "*": 2, "/": 2, "^": 3, ".": 4}
var StringPriority = 4
var NumberPriority = 4
var FunctionPriority = 4

type Node struct {
	Type          int
	Value         *Value
	Params        []*Node
	Priority      int
	TokenPosition int
}

func CreateAsOperation(operation string, params []*Node, tokenPosition int) *Node {
	node := &Node{
		Type: TypeOperation,
		Value: &Value{
			Type:      Atom,
			StringVal: &operation,
		},
		Params:        params,
		Priority:      OperationPriority[operation],
		TokenPosition: tokenPosition,
	}

	if !node.IsMathematicalOperation() {
		node.Priority = FunctionPriority
	}

	return node
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
		Priority:      NumberPriority,
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
		Priority:      StringPriority,
		TokenPosition: tokenPosition,
	}
}

func (f *Node) String(indent int) string {
	stringIndent := strings.Repeat("    ", indent)

	branches := ""

	for _, n := range f.Params {
		if n == nil {
			continue
		}
		branches = branches + fmt.Sprintf("%s└── %s", stringIndent, n.String(indent+1))
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

func (f *Node) SetOnlyChild(node *Node) {
	f.Params = []*Node{node}
}

func (f *Node) Deprioritize() {
	f.Priority = -1
}

func (f *Node) IsMathematicalOperation() bool {
	if f.Type != TypeOperation {
		return false
	}

	_, ok := OperationPriority[*f.Value.StringVal]

	return ok && *f.Value.StringVal != "."
}

func (f *Node) IsNotCallOperation() bool {
	if f.Type != TypeOperation {
		return false
	}

	return *f.Value.StringVal != "."
}

func (f *Node) IsCallOperation() bool {
	if f.Type != TypeOperation {
		return false
	}

	return *f.Value.StringVal == "."
}

func (f *Node) IsNegatable() bool {
	return f.IsFunction() || f.IsNumber()
}

func (f *Node) IsFunction() bool {
	return f.Type == TypeOperation && !f.IsMathematicalOperation() && !f.IsCallOperation()
}

func (f *Node) IsNumber() bool {
	return f.Value.IsNumber()
}

func (f *Node) IsMinusOrPlus() bool {
	return f.Value.IsMinusOrPlus()
}
