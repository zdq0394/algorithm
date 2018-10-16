package linkedlist

import "testing"

func TestSwapPairs(t *testing.T) {
	l1 := newListNodes([]int{1, 2, 3, 4, 5, 6, 7, 8}, false)
	l2 := swapPairs(l1)
	l3 := newListNodes([]int{2, 1, 4, 3, 6, 5, 8, 7}, false)
	if !equalTwoList(l2, l3) {
		t.Fail()
	}
}
