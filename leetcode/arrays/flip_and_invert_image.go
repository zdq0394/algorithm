package arrays

func flipAndInvertImage(A [][]int) [][]int {
	if A == nil || len(A) == 0 {
		return nil
	}
	for _, subA := range A {
		if subA != nil {
			var i int
			for i = 0; i < len(subA)/2; i++ {
				subA[i], subA[len(subA)-1-i] = 1-subA[len(subA)-1-i], 1-subA[i]
			}
			if len(subA)%2 != 0 {
				subA[i] = 1 - subA[i]
			}
		}
	}
	return A
	return A
}
