package p07

func numJewelsInStones(J string, S string) int {
	if len(J) == 0 {
		return 0
	}
	if len(S) == 0 {
		return 0
	}
	p := struct{}{}
	h := map[byte]struct{}{}
	for i := 0; i < len(J); i++ {
		h[J[i]] = p
	}
	c := 0
	for i := 0; i < len(S); i++ {
		if _, ok := h[S[i]]; ok {
			c++
		}
	}
	return c
}
