package leetcode

func searchMatrix(matrix [][]int, target int) bool {

	for _, row := range(matrix) {

		for _,ele := range(row) {

			if ele == target {
				return true
			}
		}

	}
	return false
    
}