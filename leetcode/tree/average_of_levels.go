package tree

func averageOfLevels(root *TreeNode) []float64 {
	ret := []float64{}
	stack := []*TreeNode{root}
	curNodesBak := 1
	curNodes := 1
	nextNodes := 0
	levelSum := 0

	for len(stack) != 0 {
		p := stack[0]
		stack = stack[1:]
		curNodes--
		levelSum += p.Val
		if p.Left != nil {
			stack = append(stack, p.Left)
			nextNodes++
		}
		if p.Right != nil {
			stack = append(stack, p.Right)
			nextNodes++
		}
		if curNodes == 0 {
			ret = append(ret, float64(levelSum)/float64(curNodesBak))
			curNodes = nextNodes
			curNodesBak = nextNodes
			nextNodes = 0
			levelSum = 0
		}
	}
	return ret
}
