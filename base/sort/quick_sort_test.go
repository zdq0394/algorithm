package sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	quickSort(arr, 0, len(arr)-1)
	if !equals(arr, result) {
		t.Fail()
	}
}
