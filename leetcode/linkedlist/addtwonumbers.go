package linkedlist

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	addition := 0
	for {
		t := p1.Val + p2.Val + addition
		p1.Val = t % 10
		addition = t / 10
		if p1.Next == nil || p2.Next == nil {
			break
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	if p1.Next == nil {
		p1.Next = p2.Next
	}
	for p1.Next != nil {
		t := p1.Next.Val + addition
		p1.Next.Val = t % 10
		addition = t / 10
		p1 = p1.Next
	}
	if addition > 0 {
		p1.Next = &ListNode{
			Val: addition,
		}
	}
	return l1
}
