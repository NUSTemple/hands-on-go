package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestIsPalindrome1(t *testing.T)  {

    s := 123
	expectedOutput := false
	funcOutput := isPalindrome(s)
	assert.Equal(t, expectedOutput, funcOutput)
}

func TestIsPalindrome2(t *testing.T)  {

    s := 121
	expectedOutput := true
	funcOutput := isPalindrome(s)
	assert.Equal(t, expectedOutput, funcOutput)
}

func TestIsPalindrome3(t *testing.T)  {

    s := 10
	expectedOutput := false
	funcOutput := isPalindrome(s)
	assert.Equal(t, expectedOutput, funcOutput)
}
