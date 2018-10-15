package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	fakeHead := &ListNode{
		Next: head,
	}
	p, q := fakeHead, head
	for p.Next != nil && q.Next != nil {
		p.Next = q.Next
		q.Next = q.Next.Next
		p.Next.Next = q
		p = q
		q = p.Next
	}
	return fakeHead.Next
}

func main() {

}
