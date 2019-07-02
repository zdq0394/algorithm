package maxpoints

func maxPoints(points [][]int) int {
	n := len(points)
	diffMap := map[Point]int{}
	for i := 0; i < n; i++ {
		p := Point{points[i][0], points[i][1]}
		diffMap[p]++
	}

	size := len(diffMap)
	if size <= 2 {
		return n
	}

	max := 0
	ppoints := make([]Point, 0, size)

	for tp := range diffMap {
		ppoints = append(ppoints, tp)
	}

	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			count := 0
			for k := 0; k < size; k++ {
				if isSameLine(ppoints[i], ppoints[j], ppoints[k]) {
					count += diffMap[ppoints[k]]
				}
			}
			if count > max {
				max = count
			}
		}
	}
	return max
}

type Point struct {
	x int
	y int
}

func isSameLine(p1, p2, p3 Point) bool {
	return (p2.y-p1.y)*(p3.x-p1.x) == (p3.y-p1.y)*(p2.x-p1.x)
}
