package max_depth

import "testing"

//       3
//      / \
//     9  20
//       /  \
//      15   7
func sampleTree() *TreeNode {
	return &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
}

func TestMaxDepth(t *testing.T) {
	if got := MaxDepth(sampleTree()); got != 3 {
		t.Errorf("MaxDepth = %d, want 3", got)
	}
	if got := MaxDepth(nil); got != 0 {
		t.Errorf("MaxDepth(nil) = %d, want 0", got)
	}
}
