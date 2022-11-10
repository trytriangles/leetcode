// You are given a perfect binary tree where all leaves are on the same level,
// and every parent has two children. The binary tree has the following
// definition:
//
//	struct Node {
//	  int val;
//	  Node *left;
//	  Node *right;
//	  Node *next;
//	}
//
// Populate each next pointer to point to its next right node. If there is no
// next right node, the next pointer should be set to NULL.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type TrackedNode struct {
	N     *Node
	Level int // Distance from the root node.
}

// connect takes the root of a binary tree, and sets each element's Next field
// to the node that would sit to the right of it on a traditional visual
// depiction of a tree.
func connect(n *Node) *Node {
	// Short circuit: a nil root requires no changes.
	if n == nil {
		return nil
	}

	// Do an iterative, queue-based breadth-first search of the tree, storing
	// each node as a TrackedNode that knows its distance from the root.
	q := []TrackedNode{{n, 0}}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		// If the next element to be processed is the distance from the root,
		// it's sitting to the right of the current element, so set the Next
		// field.
		if len(q) > 0 && q[0].Level == v.Level {
			v.N.Next = q[0].N
		}
		// Then push the left and right children onto the queue.
		if v.N.Left != nil {
			q = append(q, TrackedNode{v.N.Left, v.Level + 1})
		}
		if v.N.Right != nil {
			q = append(q, TrackedNode{v.N.Right, v.Level + 1})
		}
	}
	return n
}
