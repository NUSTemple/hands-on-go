package leetcode

func convert(s string, numRows int) string {
	// return original string if numRows is 1	
	if numRows ==  1 {
		return s
	}

	//define direction flag
	// 1 means down , -1 means up
	direction := 1

	//initialize matrix
	matrix := make([][]byte, numRows)

	//initialize row
	r := 0

	for i := range s {
		matrix[r] = append(matrix[r] , s[i])
		if r == numRows -1 {
			direction = -1
		} else if r == 0 {
			direction = 1
		}

		r = r + direction
	}

	var res []byte
	//join matrix into string
	for r := range(matrix) {
		res = append(res , matrix[r]...)
	}
	return string(res)
}
