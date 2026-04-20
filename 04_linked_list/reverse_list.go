package linkedlist_practice

// Node is a singly-linked list node.
type Node struct {
	Val  int
	Next *Node
}

// ReverseList reverses a singly-linked list and returns the new head.
func ReverseList(head *Node) *Node {
	// TODO: implement
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

// helper: collect list into slice
func toSlice(head *Node) []int {
	var out []int
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}
