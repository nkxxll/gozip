package btree

import "sort"

type Value struct {
	Rune rune
	Node bool
}

type Node struct {
	value Value
	rate  float64
	left  *Node
	right *Node
}

type BTree struct {
	head *Node
}

func (b *BTree) Traverse() string {
	return b.head.Traverse("")
}

func (b *BTree) Build(nodeList []Node) {
	sort.Slice(nodeList, func(i, j int) bool {
		return nodeList[i].rate < nodeList[j].rate
	})
	// get the last two nodes
	// add the rates
	// make a new node
	// add the last two nodes to the last node
	// add the new node to the array
	// again
}

func (n *Node) Traverse(res string) string {
	if n == nil {
		return res
	}
	if n.value.Node == false {
		res += string(n.value.Rune)
	}
	res = n.left.Traverse(res)
	res = n.right.Traverse(res)
	return res
}

func DefaultNode() Node {
	return Node{
		value: Value{Node: true},
		rate:  0,
		left:  nil,
		right: nil,
	}
}

func NewNode(value rune, rate float64, left, right *Node) Node {
	return Node{
		value: Value{Rune: value},
		rate:  rate,
		left:  left,
		right: right,
	}
}

func DefaultBTree() BTree {
	return BTree{
		head: nil,
	}
}
