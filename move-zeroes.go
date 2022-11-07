func moveZeroes(nums []int) {
	currentSlot := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[currentSlot] = nums[currentSlot], nums[i]
			currentSlot++
		}
	}
}

