package arrays

import "testing"

func TestThirdMax(t *testing.T) {
	if thirdMax([]int{1, 1, 2}) != 2 {
		t.Fail()
	}
}
