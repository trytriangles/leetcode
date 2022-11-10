// middleNode returns the middle node of a linked list. If there are two middle
// nodes (i.e. the length of the list is an even number), the second one is
// returned.
func middleNode(head *ListNode) *ListNode {
	// Short circuit: for a list of one element, the middle is that element.
	if head.Next == nil {
		return head
	}

	// Iterate the list using two pointers; oneCounter will move one step at a
	// time, twoCounter two steps at a time. When twoCounter points to the end,
	// oneCounter will point to the middle.
	oneCounter := head
	twoCounter := head
	for twoCounter != nil && twoCounter.Next != nil {
		oneCounter = oneCounter.Next
		twoCounter = twoCounter.Next.Next
	}
	return oneCounter
}
