package merge_k_sorted

// Node is a singly-linked list node.
type Node struct {
	Val  int
	Next *Node
}

// MergeKLists merges k sorted linked lists into one sorted list.
// Input: [1→4→5, 1→3→4, 2→6]  →  1→1→2→3→4→4→5→6
func MergeKLists(lists []*Node) *Node {
	// TODO: implement (divide-and-conquer merge pairs, O(n log k))
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
