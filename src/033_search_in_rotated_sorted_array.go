package leetcode


func search(nums []int, target int) int {

	first := nums[0]
	direction := 1 //1 forwards, -1 backwards
	var i int

	if first == target {
		return 0
	} else if first > target {
		i = len(nums) - 1 
		direction = -1
	} else {
		i = 0
		direction = 1
	}
	
	for nums[i] != target {

		i += direction
		if i >= len(nums) || i < 0 || (nums[i] - nums[i - direction]) * direction < 0 {
			return -1
		}
	}

    return i
}