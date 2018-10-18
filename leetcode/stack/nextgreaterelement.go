package stack

func nextGreaterElement(findNums []int, nums []int) []int {
	h := map[int]int{}
	for i, v := range nums {
		h[v] = i
	}
	ret := []int{}
	for _, v := range findNums {
		var j int
		for j = h[v] + 1; j < len(nums); j++ {
			if nums[j] > v {
				ret = append(ret, nums[j])
				break
			}
		}
		if j == len(nums) {
			ret = append(ret, -1)
		}
	}
	return ret
}
