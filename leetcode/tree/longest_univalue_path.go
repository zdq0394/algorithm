package tree

func help(n *TreeNode, max *int) int {
	if n == nil {
		return 0
	}
	l := help(n.Left, max)
	r := help(n.Right, max)
	left, right := 0, 0
	if n.Left != nil && n.Left.Val == n.Val {
		left = l + 1
	}
	if n.Right != nil && n.Right.Val == n.Val {
		right = r + 1
	}
	if left+right > *max {
		*max = left + right
	}
	if left > right {
		return left
	}
	return right
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var value int
	help(root, &value)
	return value
}
