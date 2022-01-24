package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int][]int)

	for i, s := range nums {
		if _, ok := m[s]; ok {
			var r = target - s
			if _, ok := m[r]; ok {

				return []int{i, m[r][0]}

			}
			m[s] = append(m[s], i)
		} else {

			var r = target - s
			if _, ok := m[r]; ok {

				return []int{i, m[r][0]}

			}

			m[s] = []int{i}
		}
	}
	return []int{-1, -1}
}
