package detect_cycle

// Node is a singly-linked list node.
type Node struct {
	Val  int
	Next *Node
}

// DetectCycle returns the node where the cycle begins, or nil if no cycle.
func DetectCycle(head *Node) *Node {
	// TODO: implement (Floyd's with meeting-point phase 2)
	return nil
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
