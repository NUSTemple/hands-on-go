package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTwoSum1(t *testing.T)  {
	expected_output := []int {1, 3}
	func_output := twoSum([]int {0, 1, 2, 5},  6)
	assert.True(t, unorderedEqualInt(expected_output, func_output))
}


func TestTwoSum2(t *testing.T)  {
	expected_output := []int {0, 1}
	func_output := twoSum([]int {2,7,11,15},  9)
	assert.True(t, unorderedEqualInt(expected_output, func_output))
}

func TestTwoSum3(t *testing.T)  {
	expected_output := []int {2, 1}
	func_output := twoSum([]int {3, 2, 4},  6)
	assert.True(t, unorderedEqualInt(expected_output, func_output))
}

func TestTwoSum4(t *testing.T)  {
	expected_output := []int {2, 3}
	func_output := twoSum([]int {1, 2, 3, 3, 4},  6)
	assert.True(t, unorderedEqualInt(expected_output, func_output))
}

func TestTwoSum5(t *testing.T)  {
	expected_output := []int {-1, -1}
	func_output := twoSum([]int {1, 2, 3, 3, 4},  8)
	assert.True(t, unorderedEqualInt(expected_output, func_output))
}
