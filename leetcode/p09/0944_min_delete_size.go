package p09

func minDeletionSize(A []string) int {
	rows := len(A)
	cols := len(A[0])
	c := 0
	for j := 0; j < cols; j++ {
		for i := 0; i < rows-1; i++ {
			if A[i][j] > A[i+1][j] {
				c++
				break
			}
		}
	}
	return c
}
