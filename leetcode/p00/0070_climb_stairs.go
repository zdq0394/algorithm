package p00

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	pre2 := 1
	pre1 := 2
	for i := 3; i <= n; i++ {
		pre1, pre2 = pre1+pre2, pre1
	}
	return pre1
}
