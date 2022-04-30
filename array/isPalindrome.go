package array

import (
	"strconv"
)

func IsPalindrome(x int) bool {
	str := strconv.Itoa(x)
	strLen := len(str)
	for i := 0; i <= strLen/2; i++ {
		if str[i] != str[strLen-1-i] {
			return false
		}
	}
	return true
}
