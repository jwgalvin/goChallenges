package linkedlist_practice

import (
	"reflect"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	cases := []struct {
		inputs [][]int
		want   []int
	}{
		{[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{[][]int{}, nil},
		{[][]int{{}}, nil},
	}
	for _, c := range cases {
		lists := make([]*Node, len(c.inputs))
		for i, s := range c.inputs {
			lists[i] = fromSlice(s)
		}
		got := toSlice(MergeKLists(lists))
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("MergeKLists(%v) = %v, want %v", c.inputs, got, c.want)
		}
	}
}
