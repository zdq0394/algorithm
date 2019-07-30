package p04

func findComplement(num int) int {
	numC := 1
	for numC <= num {
		numC <<= 1
	}
	numC = numC - 1
	return numC ^ num
}
