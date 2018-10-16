package linkedlist

import "testing"

func TestReverseList(t *testing.T) {
	l1 := newListNodes([]int{1, 2, 3, 4, 5, 6, 7, 8}, false)
	l2 := newListNodes([]int{8, 7, 6, 5, 4, 3, 2, 1}, false)
	l3 := reverseList(l1)
	if !equalTwoList(l2, l3) {
		t.Fail()
	}
}
