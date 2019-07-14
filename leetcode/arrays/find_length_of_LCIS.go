package arrays

func findLengthOfLCIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	count := 0
	var i int
	for i < len(nums)-1 {
		var j int
		for j = i + 1; j < len(nums); j++ {
			if nums[j] <= nums[j-1] {
				break
			}
		}
		c := j - i
		if c > count {
			count = c
		}
		i = j
	}
	return count
}
