class Solution:
    def removeDuplicates(self, nums: list[int]) -> int:
        if len(nums) == 0:
            return 0
        result = 1
        next_placement_index = 1
        for index in range(1, len(nums)):
            if nums[index] != nums[index - 1]:
                result += 1
                nums[next_placement_index] = nums[index]
                next_placement_index += 1
        return result
