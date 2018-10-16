package linkedlist

import "testing"

var l1 = newListNodes([]int{10, 20, 30, 40, 47, 81}, false)
var l2 = newListNodes([]int{11, 20, 22, 23, 60, 81}, false)
var l = newListNodes([]int{10, 11, 20, 20, 22, 23, 30, 40, 47, 60, 81, 81}, false)

func TestMergeTwoLists1(t *testing.T) {
	l3 := mergeTwoLists1(l1, l2)
	if !equalTwoList(l3, l) {
		t.Fail()
	}
}

func TestMergeTwoLists2(t *testing.T) {
	l3 := mergeTwoLists2(l1, l2)
	if !equalTwoList(l3, l) {
		t.Fail()
	}
}
