package btree

import (
	"testing"
)

func TestSimpleBuild(t *testing.T) {
	n4 := NewNode(Value{Rune: 'd'}, 0, nil, nil)
	n3 := NewNode(Value{Rune: 'c'}, 0, nil, &n4)
	n2 := NewNode(Value{Rune: 'b'}, 0, nil, nil)
	n1 := NewNode(Value{Rune: 'a'}, 0, &n2, &n3)
	b := DefaultBTree()
	b.head = &n1
	traverse_res := b.Traverse()
	traverse_exp := "<type: char, value: a, rate: 0.000000><type: char, value: b, rate: 0.000000><type: char, value: c, rate: 0.000000><type: char, value: d, rate: 0.000000>"
	if traverse_res != traverse_exp {
		t.Fatalf("Test failed exp: %s not equal res: %s", traverse_exp, traverse_res)
	}
}

func TestStringEqTraverse(t *testing.T) {
	n4 := NewNode(Value{Rune: 'd'}, 0, nil, nil)
	n3 := NewNode(Value{Rune: 'c'}, 0, nil, &n4)
	n2 := NewNode(Value{Rune: 'b'}, 0, nil, nil)
	n1 := NewNode(Value{Rune: 'a'}, 0, &n2, &n3)
	b := DefaultBTree()
	b.head = &n1
	traverse_res := b.Traverse()
	traverse_exp := b.String()
	if traverse_exp != traverse_res {
		t.Fatalf("Test failed exp: %s not equal res: %s", traverse_exp, traverse_res)
	}
}

func TestEqTreeTrivial(t *testing.T) {
	littelTree := DefaultBTree()
	node := DefaultNode()
	node2 := DefaultNode()
	littelTree.head = &node
	testTree := DefaultBTree()
	testTree.head = &node2
	if !littelTree.Eq(&testTree) {
		t.Fatal("Test failed exp: true not equal res: false")
	}
}

func TestEqTreeDepth1(t *testing.T) {
	littelTree := DefaultBTree()
	testTree := DefaultBTree()
	nodes := make([]Node, 6)
	nodes[0] = NewNode(Value{Rune: 'a'}, 0, nil, nil)
	nodes[1] = NewNode(Value{Rune: 'b'}, 0, nil, nil)
	nodes[2] = NewNode(Value{Rune: 'c'}, 0, &nodes[1], &nodes[0])
	nodes[3] = NewNode(Value{Rune: 'a'}, 0, nil, nil)
	nodes[4] = NewNode(Value{Rune: 'b'}, 0, nil, nil)
	nodes[5] = NewNode(Value{Rune: 'c'}, 0, &nodes[4], &nodes[3])
	littelTree.head = &nodes[2]
	testTree.head = &nodes[5]
	if !littelTree.Eq(&testTree) {
		t.Fatal("Test failed exp: true not equal res: false")
	}
}

func TestReduceOnePair(t *testing.T) {
	chars := make([]Node, 5)
	chars[0] = NewNode(Value{Rune: 'a'}, 50, nil, nil)
	chars[1] = NewNode(Value{Rune: 'b'}, 25, nil, nil)
	chars[2] = NewNode(Value{Rune: 'c'}, 12, nil, nil)
	chars[3] = NewNode(Value{Rune: 'd'}, 7, nil, nil)
	chars[4] = NewNode(Value{Rune: 'e'}, 6, nil, nil)
	chars = reduceOnePair(chars)
	if len(chars) >= 5 {
		t.Fatal("len too big")
	}
}

func TestSortList(t *testing.T) {
	res := ""
	exp := "abcde"

	chars := make([]Node, 5)
	chars[4] = NewNode(Value{Rune: 'a'}, 50, nil, nil)
	chars[2] = NewNode(Value{Rune: 'b'}, 25, nil, nil)
	chars[3] = NewNode(Value{Rune: 'c'}, 12, nil, nil)
	chars[1] = NewNode(Value{Rune: 'd'}, 7, nil, nil)
	chars[0] = NewNode(Value{Rune: 'e'}, 6, nil, nil)

	// this should be a mutation
	sortNodeList(chars)

	for _, el := range chars {
		res += string(el.value.Rune)
	}

	if res != exp {
		t.Fatalf("res: %s is not %s", res, exp)
	}
}

func TestBuildEndless(t *testing.T) {
	chars := make([]Node, 5)
	chars[0] = NewNode(Value{Rune: 'a'}, 50, nil, nil)
	chars[1] = NewNode(Value{Rune: 'b'}, 25, nil, nil)
	chars[2] = NewNode(Value{Rune: 'c'}, 12, nil, nil)
	chars[3] = NewNode(Value{Rune: 'd'}, 7, nil, nil)
	chars[4] = NewNode(Value{Rune: 'e'}, 6, nil, nil)
	b := DefaultBTree()
	b.Build(chars)
}

func TestBuildSimpleTree(t *testing.T) {
	chars := make([]Node, 2)
	chars[0] = NewNode(Value{Rune: 'a', Node: false}, 50, nil, nil)
	chars[1] = NewNode(Value{Rune: 'b', Node: false}, 25, nil, nil)
	b := DefaultBTree()
	b.Build(chars)
	// build e
	e := DefaultBTree()
	// make the nodes

	epxch := make([]Node, 5)
	epxch[0] = NewNode(Value{Rune: 'a', Node: false}, 50, nil, nil)
	epxch[1] = NewNode(Value{Rune: 'b', Node: false}, 25, nil, nil)
	extra := NewNode(NewValue(true, ' '), 75, &epxch[0], &epxch[1])
	e.head = &extra
	if !b.Eq(&e) {
		t.Fatalf("Btree res: %s was not equal to exp: %s", b.String(), e.String())
	}
}

func TestBuildTree(t *testing.T) {
	chars := make([]Node, 5)
	chars[0] = NewNode(Value{Rune: 'a', Node: false}, 50, nil, nil)
	chars[1] = NewNode(Value{Rune: 'b', Node: false}, 25, nil, nil)
	chars[2] = NewNode(Value{Rune: 'c', Node: false}, 12, nil, nil)
	chars[3] = NewNode(Value{Rune: 'd', Node: false}, 7, nil, nil)
	chars[4] = NewNode(Value{Rune: 'e', Node: false}, 6, nil, nil)
	b := DefaultBTree()
	b.Build(chars)
	// build e
	e := DefaultBTree()
	// make the nodes

	epxch := make([]Node, 5)
	epxch[0] = NewNode(Value{Rune: 'a', Node: false}, 50, nil, nil)
	epxch[1] = NewNode(Value{Rune: 'b', Node: false}, 25, nil, nil)
	epxch[2] = NewNode(Value{Rune: 'c', Node: false}, 12, nil, nil)
	epxch[3] = NewNode(Value{Rune: 'd', Node: false}, 7, nil, nil)
	epxch[4] = NewNode(Value{Rune: 'e', Node: false}, 6, nil, nil)
	extra := make([]Node, 5)
	extra[3] = NewNode(Value{Node: true}, 13, &epxch[3], &epxch[4])
	extra[2] = NewNode(Value{Node: true}, 25, &extra[3], &epxch[2])
	extra[1] = NewNode(Value{Node: true}, 50, &epxch[1], &extra[2])
	extra[0] = NewNode(Value{Node: true}, 100, &epxch[0], &extra[1])
	e.head = &extra[0] // some node ...
	if !b.Eq(&e) {
		t.Fatalf("Btree res: %s was not equal to exp: %s", b.String(), e.String())
	}
}
