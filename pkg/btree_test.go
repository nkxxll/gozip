package btree

import (
	"testing"
)

func TestSimpleBuild(t *testing.T) {
	n4 := NewNode('d', 0, nil, nil)
	n3 := NewNode('c', 0, nil, &n4)
	n2 := NewNode('b', 0, nil, nil)
	n1 := NewNode('a', 0, &n2, &n3)
	b := DefaultBTree()
	b.head = &n1
	traverse_res := b.Traverse()
	traverse_exp := "abcd"
	if traverse_res != traverse_exp {
		t.Fatalf("Test failed exp: %s not equal res: %s", traverse_exp, traverse_res)
	}
}
