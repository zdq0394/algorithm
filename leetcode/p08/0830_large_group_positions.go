package p08

func largeGroupPositions(S string) [][]int {
	ret := make([][]int, 0)
	curChar := S[0]
	startIndex := 0
	endIndex := 0
	for i := 1; i < len(S); i++ {
		if S[i] == curChar {
			endIndex = i
			if i == len(S)-1 {
				l := endIndex - startIndex + 1
				if l >= 3 {
					ret = append(ret, []int{startIndex, endIndex})
				}
			}
		} else {
			l := endIndex - startIndex + 1
			if l >= 3 {
				ret = append(ret, []int{startIndex, endIndex})
			}
			startIndex = i
			endIndex = i
			curChar = S[i]
		}
	}
	return ret
}
