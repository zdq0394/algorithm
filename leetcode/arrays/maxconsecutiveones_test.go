package arrays

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	nums := []int{1, 2, 1, 1, 1, 4, 5, 1, 2, 1, 1, 1, 1, 0}
	if findMaxConsecutiveOnes(nums) != 4 {
		t.Fail()
	}
}
