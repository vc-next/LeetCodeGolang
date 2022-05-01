package maxArea

import "testing"

func TestMaxArea(t *testing.T) {
	result := MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	println(result)
	if result != 49 {
		t.Errorf("result is=%d, it's invalid", result)
	}
}
