package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring1(t *testing.T) {
	input := "abcabcbb"
	funcOutput := lengthOfLongestSubstring(input)
	expectedOutput := 3
	assert.Equal(t, expectedOutput, funcOutput)
}

func TestLengthOfLongestSubstring2(t *testing.T) {
	input := "bbbbb"
	funcOutput := lengthOfLongestSubstring(input)
	expectedOutput := 1
	assert.Equal(t, expectedOutput, funcOutput)
}

func TestLengthOfLongestSubstring3(t *testing.T) {
	input := "pwwkew"
	funcOutput := lengthOfLongestSubstring(input)
	expectedOutput := 3
	assert.Equal(t, expectedOutput, funcOutput)
}

func TestLengthOfLongestSubstring4(t *testing.T) {
	input := "abba"
	funcOutput := lengthOfLongestSubstring(input)
	expectedOutput := 2
	assert.Equal(t, expectedOutput, funcOutput)
}