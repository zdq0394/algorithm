package arrays

func equalArray(a1 []int, a2 []int) bool {
	if a1 == nil && a2 == nil {
		return true
	}
	if a1 == nil || a2 == nil {
		return false
	}
	if len(a1) != len(a2) {
		return false
	}
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}
