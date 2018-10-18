package arrays

import "testing"

func TestFindDisapearedNums(t *testing.T) {
	retNums := findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	if retNums[0] != 5 || retNums[1] != 6 {
		t.Fail()
	}
}
