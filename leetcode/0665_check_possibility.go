package leetcode

func checkPossibility(nums []int) bool {
	i := 0
	j := len(nums) - 1
	for i < len(nums)-1 && nums[i] <= nums[i+1] {
		i++
	}
	if i == len(nums)-1 {
		return true
	}
	for j > 0 && nums[j] >= nums[j-1] {
		j--
	}
	if i == j-1 {
		if i == 0 || j == len(nums)-1 {
			return true
		}
		if nums[j] >= nums[i-1] {
			return true
		}
		if nums[i] <= nums[j+1] {
			return true
		}
		return false
	}
	return false

}
