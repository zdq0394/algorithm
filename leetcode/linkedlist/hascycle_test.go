package linkedlist

import (
	"testing"
)

func TestHasCycleWithUnCycledList(t *testing.T) {
	l := newListNodes([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, false)
	if hasCycle(l) {
		t.Fail()
	}
}

func TestHasCycleWithCycledList(t *testing.T) {
	l := newListNodes([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, true)
	if !hasCycle(l) {
		t.Fail()
	}
}

func TestInCircleWithCycledList(t *testing.T) {
	l := newListNodes([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, true)
	if inCircle(l) != l {
		t.Fail()
	}
}

func TestInCircleWithOutCycledList(t *testing.T) {
	l := newListNodes([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, false)
	if inCircle(l) != nil {
		t.Fail()
	}
}
