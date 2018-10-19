package linkedlist

import "testing"

func TestMiddleNode1(t *testing.T) {
	l := newListNodes([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, false)
	m := middleNode(l)
	if m.Val != 5 {
		t.Fail()
	}
}

func TestMiddleNode2(t *testing.T) {
	l := newListNodes([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, false)
	m := middleNode(l)
	if m.Val != 5 {
		t.Fail()
	}
}
