package tree

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	newFakeRoot := &TreeNode{}
	q := newFakeRoot
	stack := []*TreeNode{}
	p := root
	for p != nil || len(stack) != 0 {
		if p != nil {
			stack = append(stack, p)
			p = p.Left
		} else {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			q.Right = &TreeNode{p.Val, nil, nil}
			q = q.Right
			p = p.Right
		}
	}
	return newFakeRoot.Right
}
