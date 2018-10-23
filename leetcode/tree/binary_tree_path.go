package tree

import "fmt"

type pair struct {
	node *TreeNode
	val  int
}

func binaryTreePaths(root *TreeNode) []string {

	if root == nil {
		return []string{}
	}

	result := []string{}
	if root.Left == nil && root.Right == nil {
		return []string{fmt.Sprintf("%d", root.Val)}
	}

	if root.Left != nil {
		lPaths := binaryTreePaths(root.Left)
		for _, l := range lPaths {
			result = append(result, fmt.Sprintf("%d->%s", root.Val, l))
		}
	}
	if root.Right != nil {
		rPaths := binaryTreePaths(root.Right)
		for _, r := range rPaths {
			result = append(result, fmt.Sprintf("%d->%s", root.Val, r))
		}
	}
	return result
}
