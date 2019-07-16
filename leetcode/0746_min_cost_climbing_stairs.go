package leetcode

func minCostClimbingStairs(cost []int) int {
	a := cost[0]
	b := cost[1]
	for i := 2; i < len(cost); i++ {
		var c int
		if a < b {
			c = a + cost[i]
		} else {
			c = b + cost[i]
		}
		a = b
		b = c
	}
	if a < b {
		return a
	}
	return b
}
