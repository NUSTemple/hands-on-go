package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func getStrings() ([]string, []string) {

    a := []string{"a", "foo", "bar", "ping", "pong"}
    b := []string{"pong", "foo", "a", "bar", "ping"}

    return a, b

}

func getInt()([]int, []int) {
    a := []int{1, 2, 3, 4}
    b := []int{4, 3, 2, 1}

    return a, b
}

func TestUnorderedCompare(t *testing.T) {

	a, b := getStrings()
	assert.True(t, unorderedEqualString(a, b))
}

func TestUnorderedCompare2(t *testing.T) {

	a, b := getInt()
	assert.True(t, unorderedEqualInt(a, b))
}