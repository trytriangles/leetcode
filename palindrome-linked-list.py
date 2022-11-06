# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next


def second_half_of(node: ListNode) -> ListNode:
    slow = node
    fast = node
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
    length_is_odd = fast is not None
    if length_is_odd:
        return slow.next
    else:
        return slow


def lists_match(longer_list: ListNode, shorter_list: ListNode) -> bool:
    while shorter_list:
        if shorter_list.val != longer_list.val:
            return False
        shorter_list = shorter_list.next
        longer_list = longer_list.next
    return True


def reverse(head: ListNode):
    result = None
    current = head
    while current:
        next = current.next
        current.next = result
        result = current
        current = next
    head = result
    return head


class Solution:
    def isPalindrome(self, head: ListNode) -> bool:
        return lists_match(head, reverse(second_half_of(head)))
