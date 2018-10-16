package arrays

import "testing"

var nums = []int{2, 7, 9, 5, 6, 18}
var target = 14

func TestTwoSum1(t *testing.T) {
	a := twoSum1(nums, target)
	if !equalArray(a, []int{2, 3}) {
		t.Fail()
	}
}

func TestTwoSum2(t *testing.T) {
	a := twoSum2(nums, target)
	if !equalArray(a, []int{2, 3}) {
		t.Fail()
	}
}
