package heap

import "testing"

func TestFindKthLargest(t *testing.T) {
	k := 2
	arr := []int{3, 2, 1, 5, 6, 4}
	if findKthLargest(arr, k) != 5 {
		t.Fail()
	}
}
