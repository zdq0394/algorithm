package tree

import "testing"

func TestLowestCommonAncester(t *testing.T) {
	node1 := &TreeNode{1, nil, nil}
	node3 := &TreeNode{3, nil, nil}
	node5 := &TreeNode{5, nil, nil}
	node7 := &TreeNode{7, nil, nil}
	node2 := &TreeNode{2, node1, node3}
	node6 := &TreeNode{6, node5, node7}
	node4 := &TreeNode{4, node2, node6}
	if LowestCommonAncester(node4, node3, node5) != node4 {
		t.Fail()
	}
}
