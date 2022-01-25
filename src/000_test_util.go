package leetcode

import ("fmt")

type ListNode struct {
	Val  int
	Next *ListNode
}

func unorderedEqualString(first, second []string) bool {
	if len(first) != len(second) {
		return false
	}
	exists := make(map[string]bool)
	for _, value := range first {
		exists[value] = true
	}
	for _, value := range second {
		if !exists[value] {
			return false
		}
	}
	return true
}

func unorderedEqualInt(first, second []int) bool {
	if len(first) != len(second) {
		return false
	}
	exists := make(map[int]bool)
	for _, value := range first {
		exists[value] = true
	}
	for _, value := range second {
		if !exists[value] {
			return false
		}
	}
	return true
}

func (l ListNode) equal (l2 *ListNode) bool {
	if l.Next == nil && l2.Next == nil {
		return l.Val == l2.Val
	} else if l.Next == nil || l2.Next == nil {
		return false
	}
	return (l.Val == l2.Val) && l.Next.equal(l2.Next)
}

func initListNode(l []int) *ListNode {
	ln := new(ListNode)
	if len(l) == 0 {
		return ln
	} else if len(l) == 1 {
		ln.Val = l[0]
		return ln
	} else {
		ln.Val = l[0]
		ln.Next = initListNode(l[1:])
		return ln
	}
}

func (l ListNode) display() {
	if l == (ListNode{}) {
		fmt.Printf("Empty ListNode\n")
	} else if l.Next != nil {
		fmt.Printf("%d --> ", l.Val)
		l.Next.display()
	} else {
		fmt.Printf("%d \n", l.Val)
	}
}
