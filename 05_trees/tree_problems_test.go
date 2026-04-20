package trees_practice

import (
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {
	//     4              4
	//    / \    →       / \
	//   2   7          7   2
	//  / \ / \        / \ / \
	// 1  3 6  9      9  6 3  1
	root := &TreeNode{4,
		&TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}},
		&TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{9, nil, nil}},
	}
	inv := InvertTree(root)
	got := levelOrderVals(inv)
	want := [][]int{{4}, {7, 2}, {9, 6, 3, 1}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("InvertTree got %v, want %v", got, want)
	}
}

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

func TestBuildTree(t *testing.T) {
	root := BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	got := levelOrderVals(root)
	want := [][]int{{3}, {9, 20}, {15, 7}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("BuildTree got %v, want %v", got, want)
	}
}

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

// levelOrderVals is a test helper (BFS).
func levelOrderVals(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[i]
			level[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
		res = append(res, level)
	}
	return res
}
