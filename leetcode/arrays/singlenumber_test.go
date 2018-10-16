package arrays

import "testing"

func TestSingleNumber(t *testing.T) {
	if singleNumber([]int{2, 3, 5, 6, 3, 5, 2, 7, 6}) != 7 {
		t.Fail()
	}
}
