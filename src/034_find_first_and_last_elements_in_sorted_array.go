package leetcode

import "fmt"

func searchRange(nums []int, target int) []int {

	defaultReturn := []int {-1, -1}
	first := 0
	skipFirstCheck := false
	last := len(nums) - 1
	skipLastCheck := false

	if len(nums) < 1 {
		return defaultReturn
	} else if len(nums) == 1 {
		if nums[0] == target {
			return []int {0, 0}
		} else {
			return defaultReturn
		}
	} else if nums[first] > target || nums[last] < target {
		return defaultReturn
	} else {

		for first <  last && (!skipFirstCheck || !skipLastCheck) {
			if nums[first] == target {
				skipFirstCheck = true
			} else {
				first += 1
			}

			if nums[last] == target {
				skipLastCheck = true
			} else {
				last -= 1
			}

			fmt.Println(first, last, skipFirstCheck, skipLastCheck)
		}

		if nums[first] == target {
			return []int {first, last}
		} else {
			return defaultReturn
		}


	}

    
}