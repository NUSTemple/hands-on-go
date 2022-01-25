package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers1(t *testing.T) {
	input1 := initListNode([]int {1, 2, 3})
	input2 := initListNode([]int {4, 5, 6})
	funcOutput := addTwoNumbers(input1, input2)
	expectedOutput := initListNode([]int {5, 7, 9})
	funcOutput.display()
	expectedOutput.display()
	assert.True(t, funcOutput.equal(expectedOutput))
}


func TestAddTwoNumbers2(t *testing.T) {
	input1 := initListNode([]int {0})
	input2 := initListNode([]int {0})
	funcOutput := addTwoNumbers(input1, input2)
	expectedOutput := initListNode([]int {0})
	funcOutput.display()
	expectedOutput.display()
	assert.True(t, funcOutput.equal(expectedOutput))
}

func TestAddTwoNumbers3(t *testing.T) {
	input1 := initListNode([]int {9,9,9,9,9,9,9})
	input2 := initListNode([]int {9,9,9,9})
	funcOutput := addTwoNumbers(input1, input2)
	expectedOutput := initListNode([]int {8,9,9,9,0,0,0,1})
	funcOutput.display()
	expectedOutput.display()
	assert.True(t, funcOutput.equal(expectedOutput))
}


