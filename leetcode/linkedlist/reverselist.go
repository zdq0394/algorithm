package linkedlist

func reverseList(head *ListNode) *ListNode {
	var cur, pre *ListNode = head, nil
	for cur != nil {
		cur.Next, cur, pre = pre, cur.Next, cur
	}
	return pre
}
