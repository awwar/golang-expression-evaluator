package parser

import "slices"

type NodeList struct {
	list []*Node
	i    int
}

func (n *NodeList) Push(nodes ...*Node) {
	n.list = append(n.list, nodes...)
}

func (n *NodeList) Current() *Node {
	return n.getFromList(n.i)
}

func (n *NodeList) Left() *Node {
	return n.getFromList(n.i - 1)
}

func (n *NodeList) Right() *Node {
	return n.getFromList(n.i + 1)
}

func (n *NodeList) RightRight() *Node {
	return n.getFromList(n.i + 2)
}

func (n *NodeList) LeftLeft() *Node {
	return n.getFromList(n.i - 2)
}

func (n *NodeList) IsEnd() bool {
	return len(n.list) < 2 || n.i >= len(n.list)
}

func (n *NodeList) Next() {
	n.i = n.i + 1
}

func (n *NodeList) Rewind() {
	n.i = 0
}

func (n *NodeList) Replace(toLeft, toRight int, node *Node) {
	n.list = slices.Replace(n.list, n.i-toLeft, n.i+toRight, node)
	n.i = n.i - 1
}

func (n *NodeList) Result() []*Node {
	return n.list
}

func (n *NodeList) getFromList(i int) *Node {
	if i < 0 || i >= len(n.list) {
		return nil
	}

	return n.list[i]
}
