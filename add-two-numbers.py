# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def sumList(self, l: ListNode):
        v = l.val
        depth = 1
        n = l.next
        while n:
            v += n.val * 10**depth
            depth += 1
            n = n.next
        return v

    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        summed = self.sumList(l1) + self.sumList(l2)
        digits = [ListNode(int(c)) for c in reversed(str(summed))]
        for i in range(len(digits) - 1):
            digits[i].next = digits[i + 1]
        return digits[0]
