package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays1(t *testing.T) {
	a1 := []int {1, 3}
	a2 := []int {2}
	expectedOutput := float64(2)
	funcOutput := findMedianSortedArrays(a1, a2)
	assert.Equal(t, expectedOutput, funcOutput)
}

func TestFindMedianSortedArrays2(t *testing.T) {
	a1 := []int {1, 2}
	a2 := []int {3, 4}
	expectedOutput := float64(2.5)
	funcOutput := findMedianSortedArrays(a1, a2)
	assert.Equal(t, expectedOutput, funcOutput)
}