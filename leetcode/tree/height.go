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
