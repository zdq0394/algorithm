package p01

import (
	"testing"
)

func TestStock3(t *testing.T) {
	a := []int{3, 3, 5, 0, 0, 3, 1, 4}
	if maxProfit(a) != 6 {
		t.FailNow()
	}
}
