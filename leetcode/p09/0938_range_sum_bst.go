package p09

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	if L > root.Val {
		return rangeSumBST(root.Right, L, R)
	}

	if R < root.Val {
		return rangeSumBST(root.Left, L, R)
	}

	return rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R) + root.Val

}
