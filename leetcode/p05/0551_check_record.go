package p05

func checkRecord(s string) bool {
	aNums := 0
	lNums := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			aNums++
			if aNums > 1 {
				return false
			}
			lNums = 0
		} else if s[i] == 'L' {
			lNums++
			if lNums > 2 {
				return false
			}
		} else {
			lNums = 0
		}
	}
	return true
}
