package p09

func hasGroupsSizeX(deck []int) bool {
	h := map[int]int{}
	for i := 0; i < len(deck); i++ {
		h[deck[i]]++
	}
	minV := 10000
	for _, v := range h {
		if v < minV {
			minV = v
		}
	}
	g := minV
	for _, v := range h {
		g = gcd(g, v)
	}
	if g < 2 {
		return false
	}
	return true
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for {
		r := a % b
		if r == 0 {
			break
		}
		a = b
		b = r
	}
	return b
}
