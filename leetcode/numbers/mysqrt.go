package numbers

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	left := 0
	right := x
	var mid int
	for left < right {
		mid = left + (right-left)/2
		if x/mid >= mid {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right - 1
}
