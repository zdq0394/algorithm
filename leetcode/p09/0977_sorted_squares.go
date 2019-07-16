package p09

func sortedSquares(A []int) []int {
	if A == nil || len(A) == 0 {
		return []int{}
	}
	var i int
	for i = 0; i < len(A); i++ {
		if A[i] > 0 {
			break
		}
	}
	if i == 0 {
		for i = 0; i < len(A); i++ {
			A[i] = A[i] * A[i]
		}
		return A
	}
	if i == len(A) {
		for i = 0; i < len(A)/2; i++ {
			A[i], A[len(A)-1-i] = A[len(A)-1-i]*A[len(A)-1-i], A[i]*A[i]
		}
		if len(A)%2 == 1 {
			A[len(A)/2] = A[len(A)/2] * A[len(A)/2]
		}
		return A
	}
	r := i
	l := i - 1
	B := make([]int, 0)
	for l >= 0 && r < len(A) {
		ll := A[l] * A[l]
		rr := A[r] * A[r]
		if ll <= rr {
			B = append(B, ll)
			l--
		} else if ll > rr {
			B = append(B, rr)
			r++
		}
	}
	for l >= 0 {
		B = append(B, A[l]*A[l])
		l--
	}
	for r < len(A) {
		B = append(B, A[r]*A[r])
		r++
	}
	return B
}
