package arrays

import (
	"testing"
)

func TestIsOneBitCharacter(t *testing.T) {
	a := []int{0, 0, 1, 0}
	if isOneBitCharacter(a) {
		t.Fail()
	}
	b := []int{1, 0, 0}
	if !isOneBitCharacter(b) {
		t.Fail()
	}
}
