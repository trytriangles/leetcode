func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	resultHead := &ListNode{-1, nil}
	currentResultHead := resultHead
	for {
		if list1 == nil && list2 == nil {
			break
		}
		if list2 == nil || (list1 != nil && list1.Val < list2.Val) {
			currentResultHead.Next = list1
			list1 = list1.Next
			currentResultHead = currentResultHead.Next
			continue
		}
		currentResultHead.Next = list2
		list2 = list2.Next
		currentResultHead = currentResultHead.Next
	}

	return resultHead.Next
}