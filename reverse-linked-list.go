func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	queue := []*ListNode{}
	for head.Next != nil {
		queue = append(queue, head)
		head = head.Next
	}
	current := head
	for len(queue) > 0 {
		current.Next = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		current = current.Next
	}
	current.Next = nil
	return head
}
