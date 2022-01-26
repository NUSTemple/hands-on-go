package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestConvert1(t *testing.T) {
    s := "PAYPALISHIRING"
	expectedOutput := "PAHNAPLSIIGYIR"
	funcOutput := convert(s, 3)
	assert.Equal(t, expectedOutput, funcOutput)
}

func TestConvert2(t *testing.T) {
    s := "PAYPALISHIRING"
	expectedOutput := "PINALSIGYAHRPI"
	funcOutput := convert(s, 4)
	assert.Equal(t, expectedOutput, funcOutput)
}

func TestConvert3(t *testing.T) {
    s := "PAYPALISHIRING"
	expectedOutput := "PAYPALISHIRING"
	funcOutput := convert(s, 1)
	assert.Equal(t, expectedOutput, funcOutput)
}