package p02

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func qs(a []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	p := a[left]
	for i < j {
		for i < j && a[j] >= p {
			j--
		}
		if i < j {
			a[i] = a[j]
		}
		for i < j && a[i] <= p {
			i++
		}
		if i < j {
			a[j] = a[i]
		}
	}
	a[i] = p
	qs(a, left, i-1)
	qs(a, i+1, right)
}
