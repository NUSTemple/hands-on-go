package leetcode

// import (
// 	"fmt"
// )

func longestPalindrome(s string) string {

	//from any letter and search outwards
	var odd, even string
	res := ""

	for i := 0; i < len(s); i++ {
		
		odd = searchLongestSubstring(s, i, i)
		even = searchLongestSubstring(s, i, i+1)

		res = allLongestStrings([]string {res, odd, even})[0]
	}
	return res
}

func searchLongestSubstring(s string, l int, r int) string {
	for (l >= 0 && r < len(s) && s[l] == s[r]) {
		l -= 1
		r += 1
	}
	return s[l+1:r]
}

func allLongestStrings(inputArray []string) []string {
    max := -1 // -1 is guaranteed to be less than length of string
    var result []string
    for _, s := range inputArray {
        if len(s) < max {
            // Skip shorter string
            continue
        }
        if len(s) > max {
            // Found longer string. Update max and reset result.
            max = len(s)
            result = result[:0]
        }
        // Add to result
        result = append(result, s)
    }
    return result
}