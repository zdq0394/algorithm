package p03

func findTheDifference(s string, t string) byte {
	h := map[byte]int{}
	for i := 0; i < len(s); i++ {
		h[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		if _, ok := h[t[i]]; ok {
			h[t[i]]--
		} else {
			return t[i]
		}
	}
	for k, v := range h {
		if v != 0 {
			return k
		}
	}
	return byte(0)
}
