class Solution:
    def searchInsert(self, nums: list[int], target: int) -> int:
        for index in range(len(nums)):
            if nums[index] == target:
                return index
            elif nums[index] > target:
                return index
        return len(nums)
