package main

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

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func main() {
	l := newListNodes([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, false)
	fmt.Println(hasCycle(l))
}
