package valid_bst

import "testing"

func TestIsValidBST(t *testing.T) {
	valid := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 7},
	}
	invalid := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 6},
		Right: &TreeNode{Val: 7},
	}
	if !IsValidBST(valid) {
		t.Error("IsValidBST(valid tree) = false, want true")
	}
	if IsValidBST(invalid) {
		t.Error("IsValidBST(invalid tree) = true, want false")
	}
}
