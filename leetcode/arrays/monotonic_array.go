package arrays

func isMonotonic(A []int) bool {
	if A == nil {
		return false
	}

	if len(A) == 0 {
		return true
	}

	for i := 1; i < len(A); i++ {
		if A[i] != A[i-1] {
			if (A[i]-A[i-1])*(A[len(A)-1]-A[0]) <= 0 {
				return false
			}
		}
	}

	return true
}

func isMonotonic1(A []int) bool {
	if A == nil {
		return false
	}

	if len(A) == 0 {
		return true
	}
	var up, down int
	for i := 1; i < len(A); i++ {
		if A[i] > A[i-1] {
			up++
			if down != 0 {
				return false
			}
		} else if A[i] < A[i-1] {
			down++
			if up != 0 {
				return false
			}
		}

	}

	return true
}
