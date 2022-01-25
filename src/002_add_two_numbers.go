package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	addFlag := 0

	r := new(ListNode)

	if l1.Val+l2.Val >= 10 {

		r.Val = (l1.Val + l2.Val) % 10
		addFlag = 1

	} else {
		r.Val = l1.Val + l2.Val
	}

	if (l1.Next == nil && l2.Next == nil ) {
		if addFlag == 1 {
			r.Next = addTwoNumbersWithFlag(ListNode{}, ListNode{}, addFlag)
		} else {
			return r
		}
		
	} else if (l1.Next == nil) {
		r.Next = addTwoNumbersWithFlag(ListNode{}, *l2.Next, addFlag)
	} else if (l2.Next == nil) {
		r.Next = addTwoNumbersWithFlag(*l1.Next, ListNode{}, addFlag)
	} else {
		r.Next = addTwoNumbersWithFlag(*l1.Next, *l2.Next, addFlag)
	}

	return r
}

func addTwoNumbersWithFlag(l1 ListNode, l2 ListNode, addFlag int) *ListNode {

	r := new(ListNode)

	if (l1 == ListNode{} && l2 == ListNode{}) {
		r.Val = addFlag
		return r
	} else if l1.Val+l2.Val+addFlag >= 10 {

		r.Val = (l1.Val + l2.Val + addFlag) % 10
		addFlag = 1

	} else {
		r.Val = l1.Val + l2.Val + addFlag
		addFlag = 0
	}

	if (l1.Next == nil && l2.Next == nil ) {
		if addFlag == 1 {
			r.Next = addTwoNumbersWithFlag(ListNode{}, ListNode{}, addFlag)
		} else {
			return r
		}
		
	} else if (l1.Next == nil) {
		r.Next = addTwoNumbersWithFlag(ListNode{}, *l2.Next, addFlag)
	} else if (l2.Next == nil) {
		r.Next = addTwoNumbersWithFlag(*l1.Next, ListNode{}, addFlag)
	} else {
		r.Next = addTwoNumbersWithFlag(*l1.Next, *l2.Next, addFlag)
	}

	return r
}
