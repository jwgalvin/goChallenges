package linkedlist_practice

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	cases := []struct {
		nums []int
		n    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1}, 1, nil},
		{[]int{1, 2}, 1, []int{1}},
	}
	for _, c := range cases {
		got := toSlice(RemoveNthFromEnd(fromSlice(c.nums), c.n))
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("RemoveNthFromEnd(%v, %d) = %v, want %v", c.nums, c.n, got, c.want)
		}
	}
}
