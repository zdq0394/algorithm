package arrays

func thirdMax(nums []int) int {

	h := map[int]bool{}

	for _, v := range nums {
		h[v] = true
	}

	nums = []int{}
	for k := range h {
		nums = append(nums, k)
	}

	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}

	var big1, big2, big3 int
	big1 = nums[0]
	if nums[1] > nums[0] {
		big1 = nums[1]
		big2 = nums[0]
	} else if nums[1] < nums[0] {
		big2 = nums[1]
	}

	if nums[2] > big1 {
		big1, big2, big3 = nums[2], big1, big2
	} else if nums[2] > big2 && nums[2] < big1 {
		big2, big3 = nums[2], big2
	} else if nums[2] < big2 {
		big3 = nums[2]
	}

	for _, v := range nums[3:] {
		if v > big1 {
			big1, big2, big3 = v, big1, big2
		} else if v < big1 && v > big2 {
			big2, big3 = v, big2
		} else if v < big2 && v > big3 {
			big3 = v
		}
	}
	return big3
}
