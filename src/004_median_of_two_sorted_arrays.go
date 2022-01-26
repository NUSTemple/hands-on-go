package leetcode

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    
	mergedArray := append(nums1, nums2...)
	sort.Ints(mergedArray)

	var length int = len(mergedArray)
	if len(mergedArray) %2 == 1 {
		return float64(mergedArray[(len(mergedArray)-1)/2])

	} else {
		return float64(mergedArray[length/2] + mergedArray[length/2 - 1]) / 2.0
	}
}