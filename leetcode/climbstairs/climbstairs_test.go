package climbstairs

import "testing"

func TestClimbStairs1(t *testing.T) {
	if climbStairs1(20) != 10946 {
		t.Fail()
	}
}

func TestClimbStairs2(t *testing.T) {
	if climbStairs2(20) != 10946 {
		t.Fail()
	}
}
