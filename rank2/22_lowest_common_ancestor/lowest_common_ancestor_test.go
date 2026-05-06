package lowest_common_ancestor

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	// Tree: 3→5,1; 5→6,2; 1→0,8; 2→7,4
	n7 := &TreeNode{Val: 7}
	n4 := &TreeNode{Val: 4}
	n6 := &TreeNode{Val: 6}
	n2 := &TreeNode{2, n7, n4}
	n5 := &TreeNode{5, n6, n2}
	n0 := &TreeNode{Val: 0}
	n8 := &TreeNode{Val: 8}
	n1 := &TreeNode{1, n0, n8}
	root := &TreeNode{3, n5, n1}
	_ = root

	got := LowestCommonAncestor(root, n5, n1)
	if got == nil || got.Val != 3 {
		t.Errorf("LCA(5,1) = %v, want 3", got)
	}

	got = LowestCommonAncestor(root, n5, n4)
	if got == nil || got.Val != 5 {
		t.Errorf("LCA(5,4) = %v, want 5", got)
	}
}
