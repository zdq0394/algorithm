package tree

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 0
	}

	sum := 0
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if p.Left != nil {
			if p.Left.Left == nil && p.Left.Right == nil {
				sum += p.Left.Val
			} else {
				stack = append(stack, p.Left)
			}
		}
		if p.Right != nil {
			stack = append(stack, p.Right)
		}
	}
	return sum
}
