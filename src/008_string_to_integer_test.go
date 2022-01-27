package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestMyAtoi1(t *testing.T) {
    s := "123"
	expectedOutput := 123
	funcOutput := myAtoi(s)
	assert.Equal(t, expectedOutput, funcOutput)
}


func TestMyAtoi2(t *testing.T) {
    s := "0021"
	expectedOutput := 21
	funcOutput := myAtoi(s)
	assert.Equal(t, expectedOutput, funcOutput)
}


func TestMyAtoi3(t *testing.T) {
    s := "       0021"
	expectedOutput := 21
	funcOutput := myAtoi(s)
	assert.Equal(t, expectedOutput, funcOutput)
}

func TestMyAtoi4(t *testing.T) {
    s := "       -0021"
	expectedOutput := -21
	funcOutput := myAtoi(s)
	assert.Equal(t, expectedOutput, funcOutput)
}

func TestMyAtoi5(t *testing.T) {
    s := "       -0021 with dodf"
	expectedOutput := -21
	funcOutput := myAtoi(s)
	assert.Equal(t, expectedOutput, funcOutput)
}


func TestMyAtoi6(t *testing.T) {
    s := "       -0021 32 ith dodf"
	expectedOutput := -21
	funcOutput := myAtoi(s)
	assert.Equal(t, expectedOutput, funcOutput)
}



func TestMyAtoi7(t *testing.T) {
    s := "9223372036854775808"
	expectedOutput := 2147483647
	funcOutput := myAtoi(s)
	assert.Equal(t, expectedOutput, funcOutput)
}