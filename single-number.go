// # Problem statement
//
// You are given a non-empty array of integers, every element appears twice
// except for one. Find that single one, within linear time complexity and
// constant spatial complexity.

func singleNumber(nums []int) int {
	magic := 0
	for i := 0; i < len(nums); i++ {
		magic ^= nums[i]
	}
	return magic
}
