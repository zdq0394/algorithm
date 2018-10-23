package tree

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	min := root.Val
	min2 := 0x7FFFFFFF
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if p.Val < min {
			min = p.Val
			min2 = min
		} else if p.Val > min && p.Val < min2 {
			min2 = p.Val
		}
		if p.Val < min2 {
			if p.Right != nil {
				stack = append(stack, p.Right)
			}
			if p.Left != nil {
				stack = append(stack, p.Left)
			}
		}
	}
	if min2 == 0x7FFFFFFF {
		min2 = -1
	}
	return min2
}
