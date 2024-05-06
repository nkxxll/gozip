package btree

import (
	"fmt"
	"sort"
)

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

// todo this is not optimal string and traverse are the same
func (b *BTree) String() string {
	return b.head.Traverse("", 0)
}

func (n *Node) String() string {
	var type_ string
	if n.value.Node {
		type_ = "node"
	} else {
		type_ = "char"
	}
	var value string
	if !n.value.Node {
		value = string(n.value.Rune)
	} else {
		value = "n.a."
	}
	return fmt.Sprintf("<type: %s, value: %s, rate: %f>", type_, value, n.rate)
}

func (b *BTree) Traverse() string {
	return b.head.Traverse("", 0)
}

func (n *Node) Eq(other *Node) bool {
	// fixme: node is not defined
	if n.value.Node == true && other.value.Node == true {
		return true
	}
	if n.value.Node == true || other.value.Node == true {
		// one of the nodes is a placeholder node not a leaf
		return false
	}
	if n.value.Rune == other.value.Rune {
		return true
	}
	return false
}

func (b *BTree) Eq(other *BTree) bool {
	return b.head.recurseEq(other.head)
}

func (n *Node) recurseEq(other *Node) bool {
	res := false
	if !n.Eq(other) {
		return false
	}
	if n.left == nil && n.right == nil && other.left == nil && other.right == nil {
		return true
	}
	if n.left != nil && other.left != nil {
		res = n.left.recurseEq(other.left)
		if res == false {
			return false
		}
	}
	if n.right != nil && other.right != nil {
		res = n.right.recurseEq(other.right)
		if res == false {
			return false
		}
	}
	return res
}

func sortNodeList(nodeList []Node) {
	sort.Slice(nodeList, func(i, j int) bool {
		return nodeList[i].rate > nodeList[j].rate
	})
}

func reduceOnePair(nodeList []Node) []Node {
	var a, b Node
	a, b, nodeList = nodeList[len(nodeList)-1], nodeList[len(nodeList)-2], nodeList[:len(nodeList)-2]
	new := NewNode(Value{Node: true}, a.rate+b.rate, &b, &a)
	nodeList = append(nodeList, new)
	return nodeList
}

func (b *BTree) Build(nodeList []Node) {
	for len(nodeList) > 1 {
		sortNodeList(nodeList)
		fmt.Println(nodeList)
		nodeList = reduceOnePair(nodeList)
	}
	fmt.Println("head is" + nodeList[0].String())
	b.head = &nodeList[0]
}

func (n *Node) Traverse(res string, depth uint) string {
	if n == nil {
		return res
	}
	if n.value.Node == false {
		res += n.String()
	}
	res = n.left.Traverse(res, depth+1)
	res = n.right.Traverse(res, depth+1)
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

func NewValue(node bool, value rune) Value {
	if node == true {
		return Value{
			Node: true,
		}
	}
	return Value{
		Node: false,
		Rune: value,
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
