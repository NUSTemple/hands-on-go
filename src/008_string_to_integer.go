package leetcode

import (
	"math")

func myAtoi(s string) int {

	numMap := map[rune]int {
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}
    var nextAlphabetEsc = false
	var maxInt = int(math.Pow(2, 31))
	var res int = 0
	var sign int = 1
	for i:=0; (i < len(s)); i++  {

		if s[i] == ' ' && !nextAlphabetEsc {
			continue
		} else if s[i] == '-' && !nextAlphabetEsc {
			sign = -1
			nextAlphabetEsc = true
		} else if s[i] == '+' && !nextAlphabetEsc {
			nextAlphabetEsc = true
		} else if _, ok := numMap[rune(s[i])]; ok {
			res = res * 10 + numMap[rune(s[i])]
			nextAlphabetEsc = true
			if res >= maxInt{
				break
			}
			
		} else {
			break
		}
	}
    
    if sign > 0 {
        res = min(res, maxInt - 1)
    } else {
        res = max(-1 * res, -1 * maxInt)
    }

    return res
}