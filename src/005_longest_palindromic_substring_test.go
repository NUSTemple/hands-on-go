package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome1(t *testing.T) {
    s := "babad"
	expectedOutput := []string {"aba", "bab"}
	funcOutput := longestPalindrome(s)
	assert.Contains(t, expectedOutput, funcOutput)
}

func TestLongestPalindrome2(t *testing.T) {
    s := "cbbd"
	expectedOutput := []string {"bb"}
	funcOutput := longestPalindrome(s)
	assert.Contains(t, expectedOutput, funcOutput)
}

func TestLongestPalindrome3(t *testing.T) {
    s := "a"
	expectedOutput := []string {"a"}
	funcOutput := longestPalindrome(s)
	assert.Contains(t, expectedOutput, funcOutput)
}