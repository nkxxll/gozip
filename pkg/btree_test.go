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
	traverse_exp := "abcd"
	if traverse_res != traverse_exp {
		t.Fatalf("Test failed exp: %s not equal res: %s", traverse_exp, traverse_res)
	}
}

func TestBuildTree(t *testing.T) {
	chars := make([]Node, 5)
	chars[0] = NewNode(Value{Rune: 'a'}, 50, nil, nil)
	chars[1] = NewNode(Value{Rune: 'b'}, 25, nil, nil)
	chars[2] = NewNode(Value{Rune: 'c'}, 12, nil, nil)
	chars[3] = NewNode(Value{Rune: 'd'}, 7, nil, nil)
	chars[4] = NewNode(Value{Rune: 'e'}, 6, nil, nil)
	b := DefaultBTree()
	b.Build(chars)
	res := b.Traverse()
	exp := "abcde"
	if res != exp {
		t.Fatalf("res: %s is not %s", res, exp)
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
