// Problem statement:
//   "Given a 1-indexed array of integers numbers that is already
//    sorted in non-decreasing order, find two numbers such that
//    they add up to a specific target number."
//
// Constraints:
//   May not use the same element twice.
//   Must have O(1) space complexity.


func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
            // Return incremented indices due to the challenge
            // specifying a 1-indexed array.
			return []int{left + 1, right + 1}
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}
	// Should never reach this point under challenge parameters.
	panic("target not found")
}

