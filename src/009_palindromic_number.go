package leetcode

import (
	"strconv"
)

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	} else {
		return isPalindromeString(strconv.Itoa(x))
	}
}

func isPalindromeString(x string) bool {

	if len(x) <= 1 {
		return true 
	} else {
		return x[0] == x[len(x) - 1] && isPalindromeString(x[1:len(x)-1])
	}
}