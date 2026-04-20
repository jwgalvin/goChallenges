package trees_practice

import (
	"reflect"
	"testing"
)

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

func TestLevelOrder(t *testing.T) {
	want := [][]int{{3}, {9, 20}, {15, 7}}
	got := LevelOrder(sampleTree())
	if !reflect.DeepEqual(got, want) {
		t.Errorf("LevelOrder = %v, want %v", got, want)
	}
}
