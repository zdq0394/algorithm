package arrays

func dominantIndex(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}
	max := nums[0]
	max2 := -1
	h := map[int]int{}
	for i := 1; i < len(nums); i++ {
		h[nums[i]] = i
		if nums[i] > max {
			max2 = max
			max = nums[i]
		} else if nums[i] > max2 {
			max2 = nums[i]
		}
	}
	if max >= 2*max2 {
		return h[max]
	}
	return -1
}
