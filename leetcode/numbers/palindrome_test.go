package numbers

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !isPalindrome(121383121) {
		t.Fail()
	}
}

func TestIsNotPalindrome(t *testing.T) {
	if isPalindrome(121383120) {
		t.Fail()
	}
}
