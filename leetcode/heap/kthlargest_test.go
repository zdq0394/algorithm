package heap

import "testing"

func TestKthlargest(t *testing.T) {
	k := 3
	arr := []int{4, 5, 8, 2}
	kthLargest := Constructor(k, arr)
	if kthLargest.Add(3) != 4 {
		t.Fail()
	}
	if kthLargest.Add(5) != 5 {
		t.Fail()
	}

	if kthLargest.Add(10) != 5 {
		t.Fail()
	}
	if kthLargest.Add(9) != 8 {
		t.Fail()
	}
	if kthLargest.Add(4) != 8 {
		t.Fail()
	}
}
