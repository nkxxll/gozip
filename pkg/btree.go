package btree

import (
	"fmt"
	"slices"
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

type Byte struct {
	Val    int
	Length int
}

func reverse(current int, length int) int {
	res := 0
	for length != 0 {
		res = (res << 1) | (current & 1)
		current >>= current
		length -= 1
	}
	return res
}

func (b *BTree) Encode(chars []rune) []byte {
	cache := make(map[rune]Byte, 256)
	// this is not so performant for now
	res := []byte{}
	current := 0
	totoalLength := 0
	for _, char := range chars {
		bi, ok := cache[char]
		if !ok {
			val, length := b.FindChar(char)
			bi = Byte{Val: val, Length: length}
			cache[char] = bi
		}
		current = current<<bi.Length | bi.Val
		totoalLength += bi.Length
	}
	for current != 0 {
		res = append(res, byte(current&0xff))
		current = current >> 8
	}
	slices.Reverse(res)
	return res
}

func (n *Node) EncodeTraverse(char rune, b int) (int, bool) {
	if n.value.Node == false && n.value.Rune != char {
		return b >> 1, false
	}
	if n.value.Node == false && n.value.Rune == char {
		return b, true
	}
	b, found := n.left.EncodeTraverse(char, (b<<1)|1)
	if found {
		return b, true
	}
	b, found = n.right.EncodeTraverse(char, (b << 1))
	if found {
		return b, true
	}
	return b >> 1, false
}

func getMsb(number int) int {
	if number == 0 {
		return 0
	}
	len := 0
	for number != 0 {
		number = number >> 1
		len += 1
	}
	return len
}

func (b *BTree) FindChar(char rune) (int, int) {
	res, _ := b.head.EncodeTraverse(char, 1)
	msb := getMsb(res)
	return res & ((1 << (msb - 1)) - 1), msb - 1
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
		nodeList = reduceOnePair(nodeList)
	}
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
