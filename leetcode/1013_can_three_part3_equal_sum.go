package leetcode

func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for _, v := range A {
		sum += v
	}
	if sum%3 != 0 {
		return false
	}
	s := sum / 3
	curS := 0
	for i := 0; i < len(A); i++ {
		curS += A[i]
		if curS == s {
			curS = 0
		}
	}
	if curS != 0 {
		return false
	}
	return true
}
