package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSearchRange1(t *testing.T)  {

    s := []int {5, 7, 7, 8, 8, 10}
	target := 8
	expectedOutput := []int {3, 4}
	funcOutput := searchRange(s, target)
	assert.Equal(t, expectedOutput, funcOutput)
}


func TestSearchRange2(t *testing.T)  {

    s := []int {5, 7, 7, 8, 8, 10}
	target := 6
	expectedOutput := []int {-1, -1}
	funcOutput := searchRange(s, target)
	assert.Equal(t, expectedOutput, funcOutput)
}

func TestSearchRange3(t *testing.T)  {

    s := []int {}
	target := 0
	expectedOutput := []int {-1, -1}
	funcOutput := searchRange(s, target)
	assert.Equal(t, expectedOutput, funcOutput)
}


func TestSearchRange4(t *testing.T)  {

    s := []int {1, 3}
	target := 1
	expectedOutput := []int {0, 0}
	funcOutput := searchRange(s, target)
	assert.Equal(t, expectedOutput, funcOutput)
}



func TestSearchRange5(t *testing.T)  {

    s := []int {0,2,2,3,4,4,4}
	target := 1
	expectedOutput := []int {-1, -1}
	funcOutput := searchRange(s, target)
	assert.Equal(t, expectedOutput, funcOutput)
}

func TestSearchRange6(t *testing.T)  {

    s := []int {1, 2, 3}
	target := 2
	expectedOutput := []int {1, 1}
	funcOutput := searchRange(s, target)
	assert.Equal(t, expectedOutput, funcOutput)
}