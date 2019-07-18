package p09

func diStringMatch(S string) []int {
	d := make([]int, len(S)+1)
	i := 0
	j := len(S)
	for k := 0; k < len(S); k++ {
		if S[k] == 'I' {
			d[k] = i
			i++
		} else {
			d[k] = j
			j--
		}
	}
	d[len(S)] = j
	return d
}
