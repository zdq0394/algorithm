package stack

import (
	"strconv"
)

func calPoints(ops []string) int {
	points := []int{}
	for _, v := range ops {
		l := len(points)
		switch v {
		case "+":
			cur := points[l-2] + points[l-1]
			points = append(points, cur)
		case "D":
			cur := 2 * points[l-1]
			points = append(points, cur)
		case "C":
			points = points[:l-1]
		default:
			cur, _ := strconv.Atoi(v)
			points = append(points, cur)
		}
	}
	for i := 1; i < len(points); i++ {
		points[0] += points[i]
	}
	return points[0]
}
