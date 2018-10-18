package linkedlist

import "testing"

func TestGetIntersectionNode(t *testing.T) {
	l1 := newListNodes([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, false)
	l2 := newListNodes([]int{2, 3, 4, 5, 6, 7, 8, 9}, false)
	if getIntersectionNode(l1, l2) != nil {
		t.Fail()
	}

	l1 = newListNodes([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, false)
	l3 := newListNodes([]int{100, 101, 102}, false)
	l4 := newListNodes([]int{200, 201, 102}, false)

	p, q := l3, l4
	for p.Next != nil {
		p = p.Next
	}
	p.Next = l1

	for q.Next != nil {
		q = q.Next
	}
	q.Next = l1

	if getIntersectionNode(l3, l4).Val != 0 {
		t.Fail()
	}
}
