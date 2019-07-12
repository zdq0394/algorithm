package arrays

func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	pre2 := 0
	pre1 := 1
	cur := 0
	for i := 2; i <= N; i++ {
		cur = pre1 + pre2
		pre2 = pre1
		pre1 = cur
	}
	return cur
}
