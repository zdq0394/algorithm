package arrays

func pivotIndex(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	left := 0
	right := 0
	for i := 1; i < len(nums); i++ {
		right += nums[i]
	}

	if left == right {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		left += nums[i-1]
		right -= nums[i]
		if left == right {
			return i
		}
	}
	return -1
}
