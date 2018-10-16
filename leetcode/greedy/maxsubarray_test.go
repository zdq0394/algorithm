package greedy

import "testing"

func TestMaxSubArray(t *testing.T) {
	if 6 != maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) {
		t.Fail()
	}
}
