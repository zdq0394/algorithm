package main

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type pair struct {
	node *TreeNode
	val  int
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if isLeaf(root) {
		if root.Val == sum {
			return true
		} else {
			return false
		}
	}

	stack := list.New()
	stack.PushBack(pair{root, root.Val})
	for stack.Len() != 0 {
		e := stack.Back()
		stack.Remove(e)
		n := e.Value.(pair)
		if n.node.Left != nil {
			if isLeaf(n.node.Left) {
				if n.node.Left.Val+n.val == sum {
					return true
				}
			} else {
				stack.PushBack(pair{n.node.Left, n.node.Left.Val + n.val})
			}
		}
		if n.node.Right != nil {
			if isLeaf(n.node.Right) {
				if n.node.Right.Val+n.val == sum {
					return true
				}
			} else {
				stack.PushBack(pair{n.node.Right, n.node.Right.Val + n.val})
			}
		}
	}
	return false
}

func main() {

}
