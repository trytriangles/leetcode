package main

func search(nums []int, target int) int {
	floor := 0
	ceiling := len(nums)
	for {
		midpoint := floor + ((ceiling - floor) / 2)
		if nums[midpoint] == target {
			return midpoint
		}
		if floor == midpoint {
			break
		}
		if target < nums[midpoint] {
			ceiling = midpoint
		} else {
			floor = midpoint
		}
	}
	return -1
}
