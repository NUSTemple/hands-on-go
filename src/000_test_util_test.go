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

func TestPrintListNode(t *testing.T) {

	ln := initListNode([]int{1, 2, 3, 4})
	ln2 := initListNode([]int{})
	ln.display()
	ln2.display()
}

func TestListNode(t *testing.T) {

	ln := initListNode([]int{1, 2, 3, 4})
	ln2 := initListNode([]int{1, 2, 3, 4})
	ln3 := initListNode([]int{1, 2, 3})
	ln4 := initListNode([]int{})
	ln5 := initListNode([]int{})
	assert.True(t, ln.equal(ln2))
	assert.False(t, ln.equal(ln3))
	assert.False(t, ln.equal(ln4))
	assert.True(t, ln4.equal(ln5))
}