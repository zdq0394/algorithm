package main

import (
	"fmt"

	"github.com/zdq0394/algorithm/base/tree"
)

func main() {
	r := getCommontree()
	a := tree.InOrderTraversal(r)
	fmt.Printf("%s", "InOrder:")
	for _, v := range a {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	b := tree.PreorderTraversal(r)
	fmt.Printf("%s", "PreOrder:")
	for _, v := range b {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	c := tree.LevelOrder(r)
	fmt.Printf("%s", "LevelOrder:")
	for _, row := range c {
		for _, v := range row {
			fmt.Printf("%d ", v)
		}
	}

	fmt.Println("\nBST:")

	bst := getBStree()
	path := tree.FindPathOf(bst, 5)
	for _, p := range path {
		fmt.Printf("%d ", p.Val)
	}
}

func getCommontree() *tree.TreeNode {
	node4 := &tree.TreeNode{4, nil, nil}
	node5 := &tree.TreeNode{5, nil, nil}
	node6 := &tree.TreeNode{6, nil, nil}
	node7 := &tree.TreeNode{7, nil, nil}
	node2 := &tree.TreeNode{2, node4, node5}
	node3 := &tree.TreeNode{3, node6, node7}
	node1 := &tree.TreeNode{1, node2, node3}
	return node1
}

func getBStree() *tree.TreeNode {
	node1 := &tree.TreeNode{1, nil, nil}
	node3 := &tree.TreeNode{3, nil, nil}
	node5 := &tree.TreeNode{5, nil, nil}
	node7 := &tree.TreeNode{7, nil, nil}
	node2 := &tree.TreeNode{2, node1, node3}
	node6 := &tree.TreeNode{6, node5, node7}
	node4 := &tree.TreeNode{4, node2, node6}
	return node4
}
