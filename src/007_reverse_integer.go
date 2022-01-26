package leetcode

import (
"math")

func reverse(x int) int {

	var rem int
	var sign int
	maxInt := int(math.Pow(2, 31) - 1)

	if x == 0 {
		return x
	}

	if x > 0 {
		rem = x
		sign = 1
	} else {
		rem = -1 * x
		sign = -1
	}

	var result []int

	for rem > 0 {
		d := rem % 10
		rem = rem / 10
		result = append(result, d)
	}

	var finalResult int = 0
	for i := range result {

		finalResult = finalResult * 10 + result[i]

		if (finalResult > maxInt){
			return 0
		}

	} 

	return finalResult * sign
}