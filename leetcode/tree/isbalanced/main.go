package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func main() {
	node5 := &TreeNode{5, nil, nil}
	node4 := &TreeNode{4, nil, nil}
	node3 := &TreeNode{3, nil, nil}
	node2 := &TreeNode{2, node4, node5}
	root := &TreeNode{1, node2, node3}
}
