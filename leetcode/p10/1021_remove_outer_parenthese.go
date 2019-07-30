package p10

func removeOuterParentheses(S string) string {
	retS := make([]byte, 0)
	c := 0
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			if c > 0 {
				retS = append(retS, S[i])
			}
			c++
		} else {
			c--
			if c > 0 {
				retS = append(retS, S[i])
			}
		}
	}
	return string(retS)
}
