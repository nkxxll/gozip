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

func sortNodeList(nodeList []Node) {
	sort.Slice(nodeList, func(i, j int) bool {
		return nodeList[i].rate > nodeList[j].rate
	})
}

func reduceOnePair(nodeList []Node) {
	a, b, nodeList := nodeList[len(nodeList)-1], nodeList[len(nodeList)-2], nodeList[:len(nodeList)-2]
	new := NewNode(Value{Node: true}, a.rate+b.rate, &b, &a)
	nodeList = append(nodeList, new)
}

func (b *BTree) Build(nodeList []Node) {
	for len(nodeList) > 1 {
		sortNodeList(nodeList)
		reduceOnePair(nodeList)
	}
	b.head = &nodeList[0]
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

func NewNode(value Value, rate float64, left, right *Node) Node {
	return Node{
		value: value,
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
