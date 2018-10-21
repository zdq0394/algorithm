package tree

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	p := root
	stack := []*TreeNode{}
	preSum := 0
	for p != nil || len(stack) != 0 {
		if p != nil {
			stack = append(stack, p)
			p = p.Right
		} else {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p.Val += preSum
			preSum = p.Val
			p = p.Left
		}
	}
	return root
}
