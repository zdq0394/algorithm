package p06

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	} else {
		if t2 != nil {
			t1.Val += t2.Val
			t1.Left = mergeTrees(t1.Left, t2.Left)
			t1.Right = mergeTrees(t1.Right, t2.Right)
		}
	}
	return t1
}
