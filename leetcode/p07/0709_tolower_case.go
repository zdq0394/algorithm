package p07

func toLowerCase(str string) string {
	s := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		v := str[i]
		if v <= 'Z' && v >= 'A' {
			v = v + 'a' - 'A'
		}
		s[i] = v
	}
	return string(s)
}
