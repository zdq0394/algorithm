package p08

func transpose(A [][]int) [][]int {
	if A == nil || len(A) == 0 || len(A[0]) == 0 {
		return [][]int{}
	}
	m := len(A)
	n := len(A[0])
	newA := make([][]int, n)
	for i := 0; i < n; i++ {
		newA[i] = make([]int, m)
		for j := 0; j < m; j++ {
			newA[i][j] = A[j][i]
		}
	}
	return newA
}
