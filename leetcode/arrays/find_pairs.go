package arrays

func findPairs(nums []int, k int) int {
	if nums == nil || len(nums) == 0 || k < 0 {
		return 0
	}
	h := map[int]int{}
	count := 0
	for _, v := range nums {
		h[v]++
	}
	if k == 0 {
		for _, v := range h {
			if v > 1 {
				count++
			}
		}
		return count
	}
	hh := map[int]bool{}
	for v := range h {
		if _, ok := hh[v+k]; ok {
			count++
		}
		if _, ok := hh[v-k]; ok {
			count++
		}
		hh[v] = true
	}
	return count
}
