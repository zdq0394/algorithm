package tree

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return pathBeginNode(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func pathBeginNode(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}
	if node.Val == sum {
		return 1 + pathBeginNode(node.Left, 0) + pathBeginNode(node.Right, 0)
	}
	return pathBeginNode(node.Left, sum-node.Val) + pathBeginNode(node.Right, sum-node.Val)
}
