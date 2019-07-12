package arrays

func fairCandySwap(A []int, B []int) []int {
	ASum := 0
	AHash := map[int]bool{}
	for i := 0; i < len(A); i++ {
		ASum += A[i]
		AHash[A[i]] = true
	}
	BSum := 0
	for i := 0; i < len(B); i++ {
		BSum += B[i]
	}
	diff := ASum - BSum
	diff = diff / 2

	for j := 0; j < len(B); j++ {
		if _, ok := AHash[B[j]+diff]; ok {
			return []int{B[j] + diff, B[j]}
		}
	}

	return []int{}
}
