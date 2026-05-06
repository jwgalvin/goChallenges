package has_cycle

// Node is a singly-linked list node.
type Node struct {
	Val  int
	Next *Node
}

// HasCycle returns true if the linked list contains a cycle.
// Uses Floyd's tortoise-and-hare algorithm (O(n) time, O(1) space).
func HasCycle(head *Node) bool {
	// TODO: implement
	return false
}

// helper: build list from slice
func fromSlice(vals []int) *Node {
	dummy := &Node{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &Node{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}
