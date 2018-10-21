package tree

import "testing"

func TestPathSum(t *testing.T) {
	node5 := &TreeNode{5, nil, nil}
	node4 := &TreeNode{4, nil, nil}
	node3 := &TreeNode{3, nil, nil}
	node2 := &TreeNode{2, node4, node5}
	root := &TreeNode{1, node2, node3}
	if pathSum(root, 4) != 2 {
		t.Fail()
	}
}
