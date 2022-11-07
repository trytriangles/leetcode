func rotate(nums []int, k int) {
	lastIndex := len(nums) - 1
	temp := make([]int, len(nums))
	for i := lastIndex; i >= 0; i-- {
		newDest := i + k%len(nums)
		if newDest > lastIndex {
			newDest -= len(nums)
		}
		temp[newDest] = nums[i]
	}
	copy(nums, temp)
}

