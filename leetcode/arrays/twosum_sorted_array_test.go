package arrays

import "testing"

var sortedNums = []int{2, 5, 6, 7, 9, 18}
var sortedTarget = 14

func TestTwoSumSorted1(t *testing.T) {
	a := twoSumSorted1(sortedNums, sortedTarget)
	if !equalArray(a, []int{1, 4}) {
		t.Fail()
	}
}

func TestTwoSumSorted2(t *testing.T) {
	a := twoSumSorted2(sortedNums, sortedTarget)
	if !equalArray(a, []int{1, 4}) {
		t.Fail()
	}
}
