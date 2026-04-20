package linkedlist_practice

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	cases := []struct {
		l1, l2 []int
		want   []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{nil, nil, nil},
		{nil, []int{0}, []int{0}},
	}
	for _, c := range cases {
		got := toSlice(MergeTwoLists(fromSlice(c.l1), fromSlice(c.l2)))
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("MergeTwoLists(%v, %v) = %v, want %v", c.l1, c.l2, got, c.want)
		}
	}
}
