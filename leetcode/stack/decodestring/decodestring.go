package decodestring

func decodeString(s string) string {
	numStack := []int{}
	strStack := []string{}
	var curInt = 0
	var t = ""
	for i := 0; i < len(s); i++ {
		if isNumber(s[i]) {
			curInt = curInt*10 + int(s[i]-'0')
		} else if s[i] == '[' {
			numStack = append(numStack, curInt)
			strStack = append(strStack, t)
			curInt = 0
			t = ""
		} else if s[i] == ']' {
			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			for j := 0; j < num; j++ {
				strStack[len(strStack)-1] += t
			}
			t = strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
		} else {
			t += string(s[i])
		}
	}
	if len(strStack) == 0 {
		return t
	}
	return strStack[len(strStack)-1]
}

func isChar(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z'
}

func isNumber(b byte) bool {
	return b >= '0' && b <= '9'
}
