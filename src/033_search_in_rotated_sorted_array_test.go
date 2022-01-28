package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSearch1(t *testing.T)  {

    s := []int {4,5,6,7,0,1,2}
	target := 0
	expectedOutput := 4
	funcOutput := search(s, target)
	assert.Equal(t, expectedOutput, funcOutput)
}

func TestSearch2(t *testing.T)  {

    s := []int {4,5,6,7,0,1,2}
	target := 3
	expectedOutput := -1
	funcOutput := search(s, target)
	assert.Equal(t, expectedOutput, funcOutput)
}

func TestSearch3(t *testing.T)  {

    s := []int {1}
	target := 2
	expectedOutput := -1
	funcOutput := search(s, target)
	assert.Equal(t, expectedOutput, funcOutput)
}


func TestSearch4(t *testing.T)  {

    s := []int {1, 3}
	target := 3
	expectedOutput := 1
	funcOutput := search(s, target)
	assert.Equal(t, expectedOutput, funcOutput)
}