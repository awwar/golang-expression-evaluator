package parser

type Node interface {
	IsFilled() bool
	SetPriority(priority int)
	GetPriority() int
	String(indent int) string
	Fill(left Node, right Node)
}

//func (n *Node) IsFilled() bool {
//	return n.Value.Type != tokenizer.TypeOperation || (n.Left != nil && n.Right != nil)
//}
//
//func (n *Node) String() string {
//	return n.toStringWithIndention(0)
//}
//
//func (n *Node) toStringWithIndention(indent int) string {

//}
