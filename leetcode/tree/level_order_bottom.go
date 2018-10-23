package tree

import "container/list"

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack := list.New()
	stack.PushBack(root)
	curLevel := 1
	nextLevel := 0
	result := [][]int{}
	row := []int{}
	for stack.Len() != 0 {
		e := stack.Front()
		n := e.Value.(*TreeNode)
		stack.Remove(e)
		curLevel--

		if n.Left != nil {
			stack.PushBack(n.Left)
			nextLevel++
		}

		if n.Right != nil {
			stack.PushBack(n.Right)
			nextLevel++
		}

		row = append(row, n.Val)
		if curLevel == 0 {
			result = append([][]int{row}, result...)
			curLevel = nextLevel
			nextLevel = 0
			row = []int{}
		}
	}
	return result
}
