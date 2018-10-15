package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}

type Pair struct {
	node  *TreeNode
	depth int
}

func maxDepthIter(root *TreeNode) int {
	stack := make([]*Pair, 0)
	if root == nil {
		return 0
	}
	if root != nil {
		stack = append(stack, &Pair{root, 1})
	}
	maxDepth := 1
	for len(stack) > 0 {
		e := stack[len(stack)-1]
		if e.depth > maxDepth {
			maxDepth = e.depth
		}
		stack = stack[:len(stack)-1]
		if e.node.Left != nil {
			stack = append(stack, &Pair{e.node.Left, e.depth + 1})
		}
		if e.node.Right != nil {
			stack = append(stack, &Pair{e.node.Right, e.depth + 1})
		}
	}
	return maxDepth
}

func main() {
	fmt.Println("hello")
}
