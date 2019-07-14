package leetcode

func validMountainArray(A []int) bool {
	if A == nil {
		return false
	}
	if len(A) < 3 {
		return false
	}
	var i int
	for i = 0; i < len(A)-1; i++ {
		if A[i] >= A[i+1] {
			break
		}
	}
	if i == len(A)-1 {
		return false
	}
	if i == 0 {
		return false
	}
	for ; i < len(A)-1; i++ {
		if A[i] <= A[i+1] {
			return false
		}
	}
	return true
}
