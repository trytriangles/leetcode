class Solution:
    def countPrimes(self, n: int) -> int:
        if n < 2:
            return 0
        candidates = list(range(n))
        candidates[0] = -1
        candidates[1] = -1
        count = 0
        for x in candidates:
            if x == -1:
                continue
            count += 1
            multiples = range(x + x, n, x)
            for m in multiples:
                candidates[m] = -1
        return count
