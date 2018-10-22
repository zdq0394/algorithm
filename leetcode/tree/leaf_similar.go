package tree

func leafsOfTree(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	leaves := []int{}
	for len(stack) != 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if p.Left == nil && p.Right == nil {
			leaves = append(leaves, p.Val)
		}
		if p.Right != nil {
			stack = append(stack, p.Right)
		}
		if p.Left != nil {
			stack = append(stack, p.Left)
		}
	}
	return leaves
}
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return false
	}
	leaves1 := leafsOfTree(root1)
	leaves2 := leafsOfTree(root2)
	if len(leaves1) != len(leaves2) {
		return false
	}
	for i := 0; i < len(leaves1); i++ {
		if leaves1[i] != leaves2[i] {
			return false
		}
	}
	return true
}
