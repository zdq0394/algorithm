package linkedlist

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	l1 := newListNodes([]int{1, 2, 3, 4, 5, 6, 7}, false)
	l2 := newListNodes([]int{1, 2, 3, 4, 5, 6, 7, 9}, false)
	l3 := addTwoNumbers(l1, l2)
	l4 := newListNodes([]int{2, 4, 6, 8, 0, 3, 5, 0, 1}, false)
	if !equalTwoList(l3, l4) {
		t.Fail()
	}

}
