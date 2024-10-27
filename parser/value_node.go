package parser

type ValueNode struct {
	Value string
}

func (v *ValueNode) String(indent int) string {
	return v.Value
}

func (v *ValueNode) IsFilled() bool {
	return true
}

func (v *ValueNode) SetPriority(priority int) {
}

func (v *ValueNode) GetPriority() int {
	return 0
}
func (v *ValueNode) Fill(left, right Node) {
}
