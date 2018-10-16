package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNodes(val []int, cycle bool) *ListNode {
	if val == nil {
		return nil
	}
	l := &ListNode{
		Val: val[0],
	}
	if len(val) == 1 {
		return l
	}
	pre := l
	vRemain := val[1:]
	for i, v := range vRemain {
		q := &ListNode{
			Val: v,
		}
		pre.Next = q
		pre = pre.Next
		if i == len(vRemain)-1 && cycle {
			q.Next = l
		}
	}
	return l
}

func printListNodes(l *ListNode) {
	if l == nil {
		fmt.Println("nil list")
	}
	fmt.Println("List printing:")
	for l != nil {
		fmt.Printf("%d ", l.Val)
		l = l.Next
	}
	fmt.Println()
}

func equalTwoList(l1 *ListNode, l2 *ListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}
	for l1 != nil {
		if l2 == nil {
			return false
		}
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l2 != nil {
		return false
	}
	return true
}
