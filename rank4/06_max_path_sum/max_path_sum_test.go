package max_path_sum

import "testing"

func TestMaxPathSum(t *testing.T) {
	root := &TreeNode{-10,
		&TreeNode{9, nil, nil},
		&TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}},
	}
	got := MaxPathSum(root)
	if got != 42 {
		t.Errorf("MaxPathSum = %d, want 42", got)
	}
}
