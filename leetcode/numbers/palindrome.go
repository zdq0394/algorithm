package numbers

func isPalindrome(x int) bool {
	y := 0
	orignalX := x
	for {
		y = y*10 + x%10
		x /= 10
		if x == 0 {
			break
		}
	}
	if orignalX == y {
		return true
	}
	return false
}

func main() {
	isPalindrome(121383121)
}
