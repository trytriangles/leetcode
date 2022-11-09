// CircularQueue represents a circular queue; a wrapper around a slice of at
// most n elements, with the oldest being deleted each time a new one is
// added.
type CircularQueue struct {
	Elements      []*ListNode
	CurrentIndex  int
	PreviousIndex int
	n             int
	length        int
}

// Add places node in the queue, overwriting its oldest element.
func (cq *CircularQueue) Add(node *ListNode) {
	cq.Elements[cq.CurrentIndex] = node
	cq.PreviousIndex = cq.CurrentIndex
	cq.CurrentIndex++
	if cq.CurrentIndex > cq.length-1 {
		cq.CurrentIndex = 0
	}
}

// TargetReplacement returns the node that should replace the deletion target.
func (cq *CircularQueue) TargetReplacement() *ListNode {
	if cq.n == 1 {
		return nil
	}
	return cq.Elements[cq.newChildIndex()]
}

// Deletion target returns the node that should be deleted.
func (cq *CircularQueue) DeletionTarget() *ListNode {
	return cq.Elements[cq.deletionTargetIndex()]
}

// Deletion target returns the parent of the deletion target.
func (cq *CircularQueue) TargetParent() *ListNode {
	return cq.Elements[cq.targetParentIndex()]
}

func (cq *CircularQueue) targetParentIndex() int {
	parentIndex := cq.deletionTargetIndex() - 1
	for parentIndex < 0 {
		parentIndex = cq.length - 1
	}
	return parentIndex
}

func (cq *CircularQueue) newChildIndex() int {
	newChildIndex := cq.targetParentIndex() + 2
	for newChildIndex >= cq.length {
		newChildIndex -= cq.length
	}
	return newChildIndex
}

func (cq *CircularQueue) deletionTargetIndex() int {
	deletionTargetIndex := cq.PreviousIndex - (cq.n - 1)
	if deletionTargetIndex < 0 {
		deletionTargetIndex = cq.length + deletionTargetIndex
	}
	return deletionTargetIndex
}

// String defines the printable representation of a circular queue, being a
// multi-line string in the form
//
//	circularQueue =
//	    [0]: Val = 3
//	    [1]: Val = 5
//	    [2]: Val = 8
func (cq *CircularQueue) String() string {
	sb := &strings.Builder{}
	fmt.Fprintf(sb, "circularQueue = \n")
	for index, value := range cq.Elements {
		fmt.Fprintf(sb, "\t[%v]: Val = %v", index, value.Val)
		if index == cq.PreviousIndex {
			fmt.Fprintf(sb, "  <-- previous")
		} else if index == cq.CurrentIndex {
			fmt.Fprintf(sb, "  <-- current")
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

// NewCircularQueue returns a pointer to a new circular queue capable of
// holding up to length elements.
func NewCircularQueue(length int) *CircularQueue {
	return &CircularQueue{
		make([]*ListNode, length),
		0,
		length - 1,
		length - 1,
		length,
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Short circuit: if the list has only one element and we're asked to
	// delete element #1 from the end, that means returning a nil list.
	if n == 1 && head.Next == nil {
		return nil
	}

	// Use a circular queue to track the n+1 last elements of the list. We can
	// use this to count backwards from the end of the list once we find it,
	// saving us having to iterate the list again to find our deletion target.
	queue := NewCircularQueue(n + 1)

	// Find the end of the list, populating the circular queue as we go.
	end := head
	for end != nil {
		queue.Add(end)
		end = end.Next
	}

	// If it turns out we're deleting the very first element, we can return its
	// child right away.
	if queue.DeletionTarget() == head {
		return head.Next
	}

	queue.TargetParent().Next = queue.TargetReplacement()
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
