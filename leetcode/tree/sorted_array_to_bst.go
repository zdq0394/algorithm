package tree

func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{nums[mid], nil, nil}
	if mid > 0 {
		root.Left = sortedArrayToBST(nums[:mid])
	}
	if mid < len(nums)-1 {
		root.Right = sortedArrayToBST(nums[mid+1:])
	}
	return root
}
