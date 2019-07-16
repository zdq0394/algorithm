package p10

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	h := map[int][][]int{}
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if l, ok := h[cellDist(i, j, r0, c0)]; !ok {
				l = [][]int{}
				l = append(l, []int{i, j})
				h[cellDist(i, j, r0, c0)] = l
			} else {
				l = append(l, []int{i, j})
				h[cellDist(i, j, r0, c0)] = l
			}
		}
	}
	distList := make([]int, 0)
	for d, _ := range h {
		distList = append(distList, d)
	}
	qs(distList, 0, len(distList)-1)
	ret := [][]int{}
	for i := 0; i < len(distList); i++ {
		ret = append(ret, h[distList[i]]...)
	}
	return ret
}

func cellDist(r1, c1, r2, c2 int) int {
	r := r1 - r2
	c := c1 - c2
	sum := 0
	if r >= 0 {
		sum += r
	} else {
		sum -= r
	}
	if c >= 0 {
		sum += c
	} else {
		sum -= c
	}
	return sum
}
