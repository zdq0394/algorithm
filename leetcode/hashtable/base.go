package hashtable

func equalArrayString(a1 []string, a2 []string) bool {
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

func equalMap(m1 map[string]int, m2 map[string]int) bool {
	for k1, v1 := range m1 {
		v2, ok := m2[k1]
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}
