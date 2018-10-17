package heap

import "testing"

func TestTopKFrequent(t *testing.T) {
	arr := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	if !(arr[0] == 2 && arr[1] == 1) {
		t.Fail()
	}
}
