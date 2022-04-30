package array

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome(121) {
		t.Error("fail")
	}
}
