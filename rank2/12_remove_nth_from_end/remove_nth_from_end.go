package remove_nth_from_end

// Node is a singly-linked list node.
type Node struct {
	Val  int
	Next *Node
}

// RemoveNthFromEnd removes the nth node from the end of the list.
// Input: 1→2→3→4→5, n=2  →  1→2→3→5
func RemoveNthFromEnd(head *Node, n int) *Node {
	// TODO: implement (fast-slow pointer with dummy head, O(n))
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
