var visited = []*ListNode{}

func middleNode(head *ListNode) *ListNode {
	visited = append(visited, head)
	if head.Next == nil {
		middle := (len(visited) / 2)
		var result = visited[middle]
		visited = []*ListNode{}
		return result
	}
	return middleNode(head.Next)
}

