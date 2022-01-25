package leetcode

import "fmt"

func lengthOfLongestSubstring(s string) int {

	m := make(map[rune]int)
	var head int = 0
	var tail int = 0
	var maxLength int = 0
	var tempString string

	for i, r := range s {

		if _, ok := m[r]; ok {

			tempString = s[max(m[r]+1, head):i+1]
			head = max(m[r]+1, head)
			m[r] = i
			tail = i
			maxLength = max(maxLength, tail - head + 1)
			fmt.Println(string(r), tempString)
		} else {
			m[r] = i
			tail = i
			tempString = tempString + string(r)
			fmt.Println(string(r), tempString)
			maxLength = max(maxLength, tail - head + 1)
		}
	}
	return maxLength
    
}