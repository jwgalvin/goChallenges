package is_symmetric

import "testing"

func TestIsSymmetric(t *testing.T) {
	sym := &TreeNode{1,
		&TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}},
	}
	if !IsSymmetric(sym) {
		t.Error("IsSymmetric(symmetric tree) = false, want true")
	}

	notSym := &TreeNode{1,
		&TreeNode{2, nil, &TreeNode{3, nil, nil}},
		&TreeNode{2, nil, &TreeNode{3, nil, nil}},
	}
	if IsSymmetric(notSym) {
		t.Error("IsSymmetric(asymmetric tree) = true, want false")
	}
}
