class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        profits: List[int] = [0]
        smallest_so_far = prices[0]
        largest_so_far = prices[0]
        for price in prices[1:]:
            profits.append(price - smallest_so_far)
            if price > largest_so_far:
                largest_so_far = price
            elif price < smallest_so_far:
                smallest_so_far = price
        return max(profits)
