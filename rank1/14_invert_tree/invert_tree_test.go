package invert_tree

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
		for i := range size {
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
