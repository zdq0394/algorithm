package tree

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := height(root.Left)
	rh := height(root.Right)
	if lh > rh {
		return lh + 1
	}
	return rh + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	lh := height(root.Left)
	rh := height(root.Right)
	delta := lh - rh
	if delta > 1 || delta < -1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}
