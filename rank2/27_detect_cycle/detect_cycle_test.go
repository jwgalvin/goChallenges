package detect_cycle

import "testing"

func TestDetectCycle(t *testing.T) {
	// No cycle — should return nil
	head := fromSlice([]int{3, 2, 0, -4})
	if got := DetectCycle(head); got != nil {
		t.Errorf("DetectCycle(no cycle) = %v, want nil", got)
	}

	// Cycle: tail connects back to index 1 (Val=2)
	head = fromSlice([]int{3, 2, 0, -4})
	nodes := []*Node{}
	for n := head; n != nil; n = n.Next {
		nodes = append(nodes, n)
	}
	nodes[3].Next = nodes[1] // create cycle starting at nodes[1] (Val=2)
	got := DetectCycle(head)
	if got == nil {
		t.Fatal("DetectCycle(cycle) = nil, want cycle start node")
	}
	if got.Val != 2 {
		t.Errorf("DetectCycle(cycle) returned node with Val=%d, want 2", got.Val)
	}

	// Single-node self-loop
	single := &Node{Val: 1}
	single.Next = single
	got = DetectCycle(single)
	if got == nil || got.Val != 1 {
		t.Errorf("DetectCycle(self-loop) returned %v, want node with Val=1", got)
	}
}
