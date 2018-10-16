package linkedlist

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var newList, r *ListNode
	if l1.Val <= l2.Val {
		r = l1
		l1 = l1.Next
	} else {
		r = l2
		l2 = l2.Next
	}

	newList = r

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			r.Next = l1
			l1 = l1.Next
		} else {
			r.Next = l2
			l2 = l2.Next
		}
		r = r.Next
	}
	if l1 != nil {
		r.Next = l1
	}
	if l2 != nil {
		r.Next = l2
	}
	return newList
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l2.Next = mergeTwoLists2(l1, l2.Next)
		return l2
	}
	l1.Next = mergeTwoLists2(l1.Next, l2)
	return l1

}

func main() {
	l1 := newListNodes([]int{10, 20, 30, 40, 47, 81}, false)
	print(l1)
	l2 := newListNodes([]int{11, 20, 22, 23, 30, 50, 51, 60, 81}, false)
	print(l2)
	l := mergeTwoLists1(l1, l2)
	print(l)
}
