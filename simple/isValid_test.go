package simple

import "testing"

func TestIsValid(t *testing.T) {
	//result := IsValid("()[]{}")
	//if !result {
	//	t.Error("invalid", result)
	//}
	result2 := IsValid("){")
	if result2 {
		t.Error("invalid", result2)
	}
}
