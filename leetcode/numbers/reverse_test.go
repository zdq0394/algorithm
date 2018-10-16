package numbers

import "testing"

func TestReverse1(t *testing.T) {
	if reverse1(1456789) != 9876541 {
		t.Fail()
	}
}

func TestReverse2(t *testing.T) {
	if reverse2(1456789) != 9876541 {
		t.Fail()
	}
}
