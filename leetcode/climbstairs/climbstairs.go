package climbstairs

func climbStairs1(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	result := make([]int, n)
	result[0] = 1
	result[1] = 2

	for i := 2; i < n; i++ {
		result[i] = result[i-2] + result[i-1]
	}

	return result[n-1]
}

func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs2(n-2) + climbStairs2(n-1)
}
