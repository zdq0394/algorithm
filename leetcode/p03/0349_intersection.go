package p03

func intersection(nums1 []int, nums2 []int) []int {
	if nums1 == nil || nums2 == nil || len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	h := map[int]int{}
	for _, v := range nums1 {
		h[v] = 1
	}
	for _, v := range nums2 {
		if _, ok := h[v]; ok {
			h[v] = 2
		}
	}
	i := 0
	for k, v := range h {
		if v == 2 {
			nums1[i] = k
			i++
		}
	}
	return nums1[:i]
}
