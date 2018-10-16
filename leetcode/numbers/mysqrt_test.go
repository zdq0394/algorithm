package numbers

import "testing"

func TestMySqrt1(t *testing.T) {
	if mySqrt(10000) != 100 {
		t.Fail()
	}
}

func TestMySqrt2(t *testing.T) {
	if mySqrt(9999) != 99 {
		t.Fail()
	}
}
