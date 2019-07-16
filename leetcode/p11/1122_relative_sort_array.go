package p11

func relativeSortArray(arr1 []int, arr2 []int) []int {
	if arr2 == nil {
		return arr1
	}
	if arr1 == nil {
		return nil
	}
	h1 := map[int]int{}
	for _, v := range arr1 {
		h1[v]++
	}

	i := 0
	for _, v := range arr2 {
		times := h1[v]
		for times > 0 {
			arr1[i] = v
			i++
			times--
		}
		h1[v] = 0
	}

	start := i
	for k, v := range h1 {
		for v > 0 {
			arr1[i] = k
			i++
			v--
		}
	}
	qs(arr1, start, len(arr1)-1)
	return arr1
}
