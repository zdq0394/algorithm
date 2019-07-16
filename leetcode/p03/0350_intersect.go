package p03

func intersect(nums1 []int, nums2 []int) []int {
	if nums1 == nil || nums2 == nil || len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	h := map[int]int{}
	for _, v := range nums1 {
		h[v]++
	}
	i := 0
	for _, v := range nums2 {
		if n, ok := h[v]; ok && n > 0 {
			nums1[i] = v
			h[v]--
			i++
		}
	}
	return nums1[:i]
}
