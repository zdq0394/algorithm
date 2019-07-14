package arrays

func imageSmoother(M [][]int) [][]int {
	if M == nil || len(M) == 0 {
		return nil
	}
	m := len(M)
	newM := make([][]int, m)
	for i := 0; i < m; i++ {
		n := len(M[i])
		newN := make([]int, n)
		for j := 0; j < n; j++ {
			sum := 0
			count := 0
			for k1 := i - 1; k1 <= i+1; k1++ {
				for k2 := j - 1; k2 <= j+1; k2++ {
					if k1 >= 0 && k1 < m && k2 >= 0 && k2 < n {
						sum += M[k1][k2]
						count++
					}
				}
			}
			newN[j] = sum / count
		}
		newM[i] = newN
	}
	return newM
}
