package tree

import "testing"

func TestSumOfLeftLeaves(t *testing.T) {
	node5 := &TreeNode{5, nil, nil}
	node4 := &TreeNode{4, nil, nil}
	node3 := &TreeNode{3, nil, nil}
	node2 := &TreeNode{2, node4, node5}
	root := &TreeNode{1, node2, node3}
	if sumOfLeftLeaves(root) != 4 {
		t.Fail()
	}
}
