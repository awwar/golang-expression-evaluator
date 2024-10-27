package parser

type Node interface {
	IsFilled() bool
	SetPriority(priority int)
	GetPriority() int
	String(indent int) string
	Fill(left Node, right Node)
}
