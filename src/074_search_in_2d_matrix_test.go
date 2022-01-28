package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T)  {

    s := [][]int {{1,3,5,7},{10,11,16,20}, {23,30,34,60}}
	target := 3
	expectedOutput := true
	funcOutput := searchMatrix(s, target)
	assert.Equal(t, expectedOutput, funcOutput)
}