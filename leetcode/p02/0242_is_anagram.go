package p02

func isAnagram(s string, t string) bool {
	h1 := map[rune]int{}
	for _, r := range s {
		h1[r]++
	}
	for _, r := range t {
		h1[r]--
	}
	for _, v := range h1 {
		if v != 0 {
			return false
		}
	}
	return true
}
