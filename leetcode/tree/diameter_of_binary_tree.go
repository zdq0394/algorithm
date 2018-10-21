package tree

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_, ret := depthOfDiam(root)
	return ret
}

func depthOfDiam(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	lde, ldi := depthOfDiam(root.Left)
	rde, rdi := depthOfDiam(root.Right)
	a := lde
	if rde > a {
		a = rde
	}
	b := ldi
	if rdi > b {
		b = rdi
	}
	if lde+rde > b {
		b = lde + rde
	}
	return a + 1, b
}
