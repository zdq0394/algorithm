package p05

func reverseWords(s string) string {
	s = s + " "
	temp := make([]byte, 0)
	ret := make([]byte, len(s))
	var k int
	for j := 0; j < len(s); j++ {
		if s[j] == ' ' {
			for i := len(temp) - 1; i >= 0; i-- {
				ret[k] = temp[i]
				k++
			}
			ret[k] = ' '
			k++
			temp = make([]byte, 0)
		} else {
			temp = append(temp, s[j])
		}
	}
	return string(ret)
}
