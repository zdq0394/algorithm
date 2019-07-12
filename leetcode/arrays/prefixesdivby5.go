package arrays

func prefixesDivBy5(A []int) []bool {
	r := []bool{}
	cur := 0
	for i := 0; i < len(A); i++ {
		cur = cur*2 + A[i]
		if cur%5 == 0 {
			r = append(r, true)
		} else {
			r = append(r, false)
		}
		cur = cur % 5
	}
	return r
}
