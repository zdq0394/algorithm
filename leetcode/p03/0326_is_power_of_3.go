package p03

func isPowerOfThree1(n int) bool {
	if n == 1 {
		return true
	}
	k := 3
	for k <= n {
		if k == n {
			return true
		}
		k = k * 3
	}
	return false
}

func isPowerOfThree2(n int) bool {
	if n <= 0 {
		return false
	}

	for n%3 == 0 {
		n = n / 3
	}
	return n == 1
}
