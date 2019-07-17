package p11

func defangIPaddr(address string) string {
	s := make([]rune, len(address)+6)
	var i int
	for _, r := range address {
		if r == '.' {
			s[i] = '['
			s[i+1] = '.'
			s[i+2] = ']'
			i += 3
		} else {
			s[i] = r
			i++
		}
	}
	return string(s)
}
