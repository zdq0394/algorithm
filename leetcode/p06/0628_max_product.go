package p06

func maximumProduct(nums []int) int {
	a1, a2, a3 := -1001, -1001, -1001
	b1, b2, b3 := 1001, 1001, 1001
	for _, v := range nums {
		if v > a1 {
			a3 = a2
			a2 = a1
			a1 = v
		} else if v <= a1 && v > a2 {
			a3 = a2
			a2 = v
		} else if v <= a2 && v > a3 {
			a3 = v
		}
		if v < b1 {
			b3 = b2
			b2 = b1
			b1 = v
		} else if v >= b1 && v < b2 {
			b3 = b2
			b2 = v
		} else if v >= b2 && v < b3 {
			b3 = v
		}
	}

	c1 := b1 * b2 * a1
	c2 := a3 * a2 * a1
	if c1 > c2 {
		return c1
	}
	return c2

}
