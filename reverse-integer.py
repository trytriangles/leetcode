class Solution:
    def reverse(self, x: int) -> int:
        if x >= 0:
            s = int(str(x)[::-1])
        else:
            s = -int(str(abs(x))[::-1])
        if -(2**31) <= s < 2**31:
            return s
        return 0
