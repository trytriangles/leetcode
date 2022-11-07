func sortedSquares(nums []int) []int {
	for index, value := range nums {
		nums[index] = value * value
	}
	sort.Ints(nums)
	return nums
}
