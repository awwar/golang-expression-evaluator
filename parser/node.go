package parser

import (
	"fmt"
	"strconv"
	"strings"

	"expression_parser/program"
	"expression_parser/utility"
)

const (
	TypeOperation               int = iota
	TypeConstant                int = iota
	TypeVariable                int = iota
	TypeFlowLink                int = iota
	TypeFlowDeclaration         int = iota
	TypeFlowBranchesDeclaration int = iota
	TypeProgram                 int = iota
)

var OperationPriority = map[string]int{"+": 1, "-": 2, "*": 2, "/": 2, ">": 2, "<": 2, "=": 2, "^": 3, ".": 4}

type Node struct {
	Type          int
	Value         program.Value
	Params        []*Node
	Priority      int
	TokenPosition int
}

func CreateAsProgram(params []*Node) *Node {
	return &Node{
		Type:          TypeProgram,
		Value:         program.NewString("root"),
		Params:        params,
		Priority:      4,
		TokenPosition: 0,
	}
}

func CreateAsConstant(value string, tokenPosition int) *Node {
	return &Node{
		Type:          TypeConstant,
		Value:         program.NewString(value),
		Params:        make([]*Node, 0),
		Priority:      4,
		TokenPosition: tokenPosition,
	}
}

func CreateAsOperation(operation string, params []*Node, tokenPosition int) *Node {
	node := &Node{
		Type:          TypeOperation,
		Value:         program.NewString(operation),
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
	var valueObject program.Value

	if strings.Contains(value, ".") {
		val := utility.Must(strconv.ParseFloat(value, 64))

		valueObject = program.NewFloat(val)
	} else {
		val := utility.Must(strconv.ParseInt(value, 0, 64))

		valueObject = program.NewInteger(val)
	}

	return &Node{
		Type:          TypeConstant,
		Value:         valueObject,
		Params:        make([]*Node, 0),
		Priority:      4,
		TokenPosition: tokenPosition,
	}
}

func CreateAsString(value string, tokenPosition int) *Node {
	return &Node{
		Type:          TypeConstant,
		Value:         program.NewString(value),
		Params:        make([]*Node, 0),
		Priority:      4,
		TokenPosition: tokenPosition,
	}
}

func CreateAsFlowDeclaration(value string, params []*Node, tokenPosition int) *Node {
	return &Node{
		Type:          TypeFlowDeclaration,
		Value:         program.NewString(value),
		Params:        params,
		Priority:      4,
		TokenPosition: tokenPosition,
	}
}

func CreateAsFlowBranchesDeclaration(value string, params []*Node, tokenPosition int) *Node {
	return &Node{
		Type:          TypeFlowBranchesDeclaration,
		Value:         program.NewString(value),
		Params:        params,
		Priority:      4,
		TokenPosition: tokenPosition,
	}
}

func CreateAsFlowLink(value string, tokenPosition int) *Node {
	return &Node{
		Type:          TypeFlowLink,
		Value:         program.NewString(value),
		Params:        make([]*Node, 0),
		Priority:      4,
		TokenPosition: tokenPosition,
	}
}

func CreateAsVariable(value string, tokenPosition int) *Node {
	return &Node{
		Type:          TypeVariable,
		Value:         program.NewString(value),
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

	v, err := f.Value.ToString()
	if err != nil {
		return false
	}

	_, ok := OperationPriority[string(*v)]

	return ok && string(*v) != "."
}

func (f *Node) IsNotCallOperation() bool {
	if f.Type != TypeOperation {
		return false
	}

	v, err := f.Value.ToString()
	if err != nil {
		return false
	}

	return string(*v) != "."
}

func (f *Node) IsCallOperation() bool {
	if f.Type != TypeOperation {
		return false
	}

	v, err := f.Value.ToString()
	if err != nil {
		return false
	}

	return string(*v) == "."
}

func (f *Node) IsNegatable() bool {
	return f.IsFunction() || f.IsNumber()
}

func (f *Node) IsFunction() bool {
	return f.Type == TypeOperation && !f.IsMathematicalOperation() && !f.IsCallOperation()
}

func (f *Node) IsNumber() bool {
	return program.IsNumber(f.Value)
}

func (f *Node) IsMinusOrPlus() bool {
	return program.IsMinusOrPlus(f.Value)
}

func (f *Node) IsFlowLink() bool {
	return f.Type == TypeFlowLink
}
