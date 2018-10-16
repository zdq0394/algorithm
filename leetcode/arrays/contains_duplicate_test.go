package arrays

import "testing"

var array_nums = []int{1, 2, 3, 5, 4, 3}

func TestContainsDuplicate(t *testing.T) {
	if !containsDuplicate(array_nums) {
		t.Fail()
	}
}

func TestContainsNearbyDuplicate(t *testing.T) {
	if !containsNearbyDuplicate(array_nums, 4) {
		t.Fail()
	}
}
