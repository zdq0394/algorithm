package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var cur, pre *ListNode = head, nil
	for cur != nil {
		cur.Next, cur, pre = pre, cur.Next, cur
	}
	return pre
}

func main() {

}
