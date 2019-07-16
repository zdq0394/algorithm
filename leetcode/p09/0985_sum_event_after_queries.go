package p09

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	sum := 0
	for _, v := range A {
		if v%2 == 0 {
			sum += v
		}
	}
	l := len(queries)
	ret := make([]int, l)
	for i, q := range queries {
		if A[q[1]]%2 == 0 {
			sum -= A[q[1]]
		}
		A[q[1]] += q[0]
		if A[q[1]]%2 == 0 {
			sum += A[q[1]]
		}
		ret[i] = sum
	}
	return ret
}
