package arrays

func heightChecker(heights []int) int {
	if heights == nil || len(heights) == 0 {
		return 0
	}
	rep := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		rep[i] = heights[i]
	}
	quickSorts(rep, 0, len(rep)-1)
	var count int
	for i := 0; i < len(heights); i++ {
		if rep[i] != heights[i] {
			count++
		}
	}
	return count
}

func quickSorts(a []int, l, r int) {
	if l >= r {
		return
	}
	p := a[l]
	i, j := l, r
	for i < j {
		for i < j && a[j] >= p {
			j--
		}
		if i < j {
			a[i] = a[j]
		}
		for i < j && a[i] <= p {
			i++
		}
		if i < j {
			a[j] = a[i]
		}
	}
	a[i] = p
	quickSorts(a, l, i-1)
	quickSorts(a, i+1, r)
}
