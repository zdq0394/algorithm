package p08

import "fmt"

func peakIndexInMountainArray1(A []int) int {
	for i := 0; i < len(A)-1; i++ {
		if A[i+1] < A[i] {
			return i
		}
	}
	return -1
}

func peakIndexInMountainArray2(A []int) int {
	return peakIndexInArray(A, 0, len(A)-1)
}

func peakIndexInArray(A []int, l, r int) int {
	if r-l+1 < 3 {
		return -1
	}
	k := (l + r) / 2
	if (A[k] > A[k-1]) && (A[k] > A[k+1]) {
		fmt.Println("helo")
		return k
	}

	if A[k+1] > A[k] {
		return peakIndexInArray(A, k, r)
	} else {
		return peakIndexInArray(A, l, k)
	}
}
