package sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var arr = []int{2, 7, 11, 15}
	index := TwoSum(arr, 9)
	expected := []int{0, 1}
	if !reflect.DeepEqual(index, expected) {
		t.Error("Fail")
	}
}
