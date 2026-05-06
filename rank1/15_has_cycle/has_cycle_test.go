package has_cycle

import "testing"

func TestHasCycle(t *testing.T) {
	// No cycle
	head := fromSlice([]int{3, 2, 0, -4})
	if HasCycle(head) {
		t.Error("HasCycle(no cycle) = true, want false")
	}

	// Cycle: tail connects back to index 1
	head = fromSlice([]int{3, 2, 0, -4})
	nodes := []*Node{}
	for n := head; n != nil; n = n.Next {
		nodes = append(nodes, n)
	}
	nodes[3].Next = nodes[1] // create cycle
	if !HasCycle(head) {
		t.Error("HasCycle(cycle) = false, want true")
	}
}
