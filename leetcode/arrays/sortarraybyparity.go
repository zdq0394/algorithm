package arrays

func sortArrayByParity(A []int) []int {
	p, q := 0, len(A)-1
	for p < q {
		if A[p]%2 == 0 {
			p++
		} else {
			if A[q]%2 == 0 {
				A[p], A[q] = A[q], A[p]
			} else {
				q--
			}
		}
	}
	return A
}
