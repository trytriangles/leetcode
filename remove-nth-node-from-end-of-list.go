func removeNthFromEnd(head *ListNode, n int) *ListNode {
	end := head
	length := 1
	for end.Next != nil {
		length++
		end = end.Next
	}
	desiredIndex := length - n // Index of target's parent
	if desiredIndex == 0 {
		if head.Next == nil {
			return nil
		}
		return head.Next
	}
	currentIndex := 0
	end = head
	for currentIndex < desiredIndex-1 {
		end = end.Next
		currentIndex++
	}
	if end.Next == nil {
		panic("Uh oh")
	} else {
		end.Next = end.Next.Next
	}
	return head
}
