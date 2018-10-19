package tree

func FindPathOf(root *TreeNode, k int) []*TreeNode {
	if root == nil {
		return nil
	}
	path := []*TreeNode{root}
	nodes := []*TreeNode{root}
	for len(nodes) != 0 {
		p := nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		if k == p.Val {
			return path
		}
		if k < p.Val && p.Left != nil {
			nodes = append(nodes, p.Left)
			path = append(path, p.Left)
		} else if k > p.Val && p.Right != nil {
			nodes = append(nodes, p.Right)
			path = append(path, p.Right)
		}
	}
	return []*TreeNode{}
}

func LowestCommonAncester(root *TreeNode, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for root != nil {
		if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else {
			return root
		}
	}
	return nil
}
