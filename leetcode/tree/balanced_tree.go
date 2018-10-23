package tree

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
