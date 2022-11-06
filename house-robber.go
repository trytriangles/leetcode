func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	var results = make([]int, len(nums))
	results[0] = nums[0]
	results[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		results[i] = max(results[i-1], results[i-2]+nums[i])
	}
	return results[len(results)-1]
}
