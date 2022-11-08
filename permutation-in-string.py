from collections import Counter


class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        chars_remaining = Counter(s1)
        left = -1
        for index, char in enumerate(s2):
            if char in s1:
                if left == -1:
                    left = index
                if chars_remaining[char] > 0:
                    chars_remaining[char] -= 1
                    if chars_remaining.total() == 0:
                        return True
                else:
                    while s2[left] != char:
                        chars_remaining[s2[left]] += 1
                        left += 1
                    left += 1
            else:
                if left != -1:
                    chars_remaining = Counter(s1)
                    left = -1
                continue
        return False

