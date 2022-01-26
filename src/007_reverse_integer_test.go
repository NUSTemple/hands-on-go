package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestReverse1(t *testing.T) {
    s := 123
	expectedOutput := 321
	funcOutput := reverse(s)
	assert.Equal(t, expectedOutput, funcOutput)
}

func TestReverse2(t *testing.T) {
    s := -123
	expectedOutput := -321
	funcOutput := reverse(s)
	assert.Equal(t, expectedOutput, funcOutput)
}

func TestReverse3(t *testing.T) {
    s := 120
	expectedOutput := 21
	funcOutput := reverse(s)
	assert.Equal(t, expectedOutput, funcOutput)
}

func TestReverse4(t *testing.T) {
    s := 1534236469
	expectedOutput := 0
	funcOutput := reverse(s)
	assert.Equal(t, expectedOutput, funcOutput)
}