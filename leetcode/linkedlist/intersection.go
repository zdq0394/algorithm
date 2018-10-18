package linkedlist

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p := headA
	for p.Next != nil {
		p = p.Next
	}
	p.Next = headA
	return inCircle(headB)
}
