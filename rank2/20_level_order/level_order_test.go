package level_order

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

func TestLevelOrder(t *testing.T) {
	want := [][]int{{3}, {9, 20}, {15, 7}}
	got := LevelOrder(sampleTree())
	if !reflect.DeepEqual(got, want) {
		t.Errorf("LevelOrder = %v, want %v", got, want)
	}
}
