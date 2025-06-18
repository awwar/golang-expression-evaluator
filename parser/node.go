package parser

import (
	"fmt"
	"strconv"
	"strings"

	"expression_parser/program"
	"expression_parser/utility"
)

const (
	TypeOperation       int = iota
	TypeConstant        int = iota
	TypeVariable        int = iota
	TypeFlowLink        int = iota
	TypeFlowDeclaration int = iota
	TypeFlowMetadata    int = iota
	TypeProgram         int = iota
)

var OperationPriority = map[string]int{"+": 1, "-": 2, "*": 2, "/": 2, ">": 2, "<": 2, "=": 2, "^": 3, ".": 4}

type Node struct {
	Type          int
	Value         *program.Value
	Params        []*Node
	Priority      int
	TokenPosition int
}

func CreateAsProgram(params []*Node) *Node {
	return &Node{
		Type: TypeProgram,
		Value: &program.Value{
			ValueType: program.Atom,
			StringVal: utility.AsPtr("root"),
		},
		Params:        params,
		Priority:      4,
		TokenPosition: 0,
	}
}

func CreateAsConstant(value string, tokenPosition int) *Node {
	return &Node{
		Type: TypeConstant,
		Value: &program.Value{
			ValueType: program.Atom,
			StringVal: &value,
		},
		Params:        make([]*Node, 0),
		Priority:      4,
		TokenPosition: tokenPosition,
	}
}

func CreateAsOperation(operation string, params []*Node, tokenPosition int) *Node {
	node := &Node{
		Type: TypeOperation,
		Value: &program.Value{
			ValueType: program.Atom,
			StringVal: &operation,
		},
		Params:        params,
		Priority:      OperationPriority[operation],
		TokenPosition: tokenPosition,
	}

	if !node.IsMathematicalOperation() {
		node.Priority = 4
	}

	return node
}

func CreateAsNumber(value string, tokenPosition int) *Node {
	valueObject := program.Value{}

	if strings.Contains(value, ".") {
		val := utility.Must(strconv.ParseFloat(value, 64))

		valueObject.ValueType = program.Float
		valueObject.FloatVal = &val
	} else {
		val := utility.Must(strconv.ParseInt(value, 0, 64))

		valueObject.ValueType = program.Integer
		valueObject.IntVal = &val
	}

	return &Node{
		Type:          TypeConstant,
		Value:         &valueObject,
		Params:        make([]*Node, 0),
		Priority:      4,
		TokenPosition: tokenPosition,
	}
}

func CreateAsString(value string, tokenPosition int) *Node {
	return &Node{
		Type: TypeConstant,
		Value: &program.Value{
			ValueType: program.String,
			StringVal: &value,
		},
		Params:        make([]*Node, 0),
		Priority:      4,
		TokenPosition: tokenPosition,
	}
}

func CreateAsFlowDeclaration(value string, params []*Node, tokenPosition int) *Node {
	return &Node{
		Type: TypeFlowDeclaration,
		Value: &program.Value{
			ValueType: program.Atom,
			StringVal: &value,
		},
		Params:        params,
		Priority:      4,
		TokenPosition: tokenPosition,
	}
}

func CreateAsFlowMetadata(name *Node, value *Node, tokenPosition int) *Node {
	return &Node{
		Type: TypeFlowMetadata,
		Value: &program.Value{
			ValueType: program.Atom,
			StringVal: utility.AsPtr(":META"),
		},
		Params:        []*Node{name, value},
		Priority:      4,
		TokenPosition: tokenPosition,
	}
}

func CreateAsFlowLink(value string, tokenPosition int) *Node {
	return &Node{
		Type: TypeFlowLink,
		Value: &program.Value{
			ValueType: program.Atom,
			StringVal: &value,
		},
		Params:        make([]*Node, 0),
		Priority:      4,
		TokenPosition: tokenPosition,
	}
}

func CreateAsVariable(value string, tokenPosition int) *Node {
	return &Node{
		Type: TypeVariable,
		Value: &program.Value{
			ValueType: program.Atom,
			StringVal: &value,
		},
		Params:        make([]*Node, 0),
		Priority:      4,
		TokenPosition: tokenPosition,
	}
}

func (f *Node) String(indent int) string {
	if f == nil {
		return "nil"
	}

	stringIndent := strings.Repeat("    ", indent)

	branches := ""

	for _, n := range f.Params {
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

func (f *Node) IsFlowLink() bool {
	return f.Type == TypeFlowLink
}
