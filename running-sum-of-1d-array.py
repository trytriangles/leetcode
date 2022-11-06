class Solution:
    def runningSum(self, nums: List[int]) -> List[int]:
        sum = 0
        result = []
        for n in range(len(nums)):
            sum += nums[n]
            result.append(sum)
        return result
