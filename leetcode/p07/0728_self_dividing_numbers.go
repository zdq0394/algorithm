package p07

func selfDividingNumbers(left int, right int) []int {
	ret := make([]int, 0)
	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			ret = append(ret, i)
		}
	}
	return ret
}

func isSelfDividing(n int) bool {
	m := n
	for n != 0 {
		d := n % 10
		if d == 0 {
			return false
		}
		if m%d != 0 {
			return false
		}
		n = n / 10
	}
	return true
}
