package p03

func isPowerOfFour(num int) bool {
	if num <= 0 {
		return false
	}
	if num == 1 {
		return true
	}
	k := 4
	for k <= num {
		if k == num {
			return true
		}
		k = 4 * k
	}
	return false
}
