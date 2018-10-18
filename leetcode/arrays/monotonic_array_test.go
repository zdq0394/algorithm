package arrays

import "testing"

func TestIsMonotonic(t *testing.T) {
	if !isMonotonic([]int{1, 1, 2, 3, 4, 5, 7, 7, 7}) {
		t.Fail()
	}
}

func TestIsMonotonic1(t *testing.T) {
	if isMonotonic([]int{2, 1, 2, 3, 4, 5, 7, 7, 7}) {
		t.Fail()
	}
}
