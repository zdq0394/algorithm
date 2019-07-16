package p09

func largestPerimeter(A []int) int {
	qs(A, 0, len(A)-1)
	for i := len(A) - 1; i >= 2; i-- {
		if A[i-2]+A[i-1] > A[i] {
			return A[i-2] + A[i-1] + A[i]
		}
	}
	return 0
}
