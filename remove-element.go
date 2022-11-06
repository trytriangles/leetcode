func deleteElement(arr []int, start int, length int) {
	for i := start; i < len(arr)-length; i++ {
		arr[i] = arr[i+length]
	}
}

func removeElement(nums []int, val int) int {
	indexesToRemove := []int{}
	for index, num := range nums {
		if num == val {
			indexesToRemove = append(indexesToRemove, index)
		}
	}
	for offset, index := range indexesToRemove {
		deleteElement(nums, index-offset, 1)
	}
	return len(nums) - len(indexesToRemove)
}