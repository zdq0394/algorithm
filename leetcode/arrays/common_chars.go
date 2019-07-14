package arrays

func commonChars(A []string) []string {
	if A == nil || len(A) == 0 {
		return nil
	}
	existing := map[rune]int{}
	for _, c := range A[0] {
		existing[c]++
	}
	for i := 1; i < len(A); i++ {
		cur := map[rune]int{}
		for _, c := range A[i] {
			if _, ok := existing[c]; ok {
				cur[c]++
			}
		}
		for k, v := range cur {
			if existing[k] < v {
				cur[k] = existing[k]
			}
		}
		existing = cur
	}
	ret := []string{}
	for k, v := range existing {
		for i := 0; i < v; i++ {
			ret = append(ret, string(k))
		}
	}
	return ret
}
