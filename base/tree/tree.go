package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	ret := []int{}
	if root == nil {
		return []int{}
	}
	stack = append(stack, root)
	for len(stack) != 0 {
		p := stack[len(stack)-1]
		ret = append(ret, p.Val)
		stack = stack[:len(stack)-1]
		if p.Right != nil {
			stack = append(stack, p.Right)
		}
		if p.Left != nil {
			stack = append(stack, p.Left)
		}
	}
	return ret
}

func InOrderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	ret := []int{}
	p := root
	for p != nil || len(stack) != 0 {
		if p != nil {
			stack = append(stack, p)
			p = p.Left
		} else {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret = append(ret, p.Val)
			p = p.Right
		}
	}
	return ret
}

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	curLevelNum := 1
	nextLevelNum := 0
	queue := []*TreeNode{root}
	ret := [][]int{}
	curRow := []int{}
	for len(queue) != 0 {
		p := queue[0]
		curRow = append(curRow, p.Val)
		queue = queue[1:]
		curLevelNum--
		if p.Left != nil {
			queue = append(queue, p.Left)
			nextLevelNum++
		}
		if p.Right != nil {
			queue = append(queue, p.Right)
			nextLevelNum++
		}
		if curLevelNum == 0 {
			curLevelNum = nextLevelNum
			nextLevelNum = 0
			ret = append(ret, curRow)
			curRow = []int{}
		}
	}
	return ret
}
