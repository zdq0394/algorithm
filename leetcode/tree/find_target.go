package tree

func findTarget(root *TreeNode, k int) bool {
	h := map[int]*TreeNode{}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if v, ok := h[k-p.Val]; ok && p != v {
			return true
		}
		h[p.Val] = p
		if p.Right != nil {
			stack = append(stack, p.Right)
		}
		if p.Left != nil {
			stack = append(stack, p.Left)
		}
	}
	return false
}
