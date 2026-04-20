package linkedlist_practice

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	cases := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1}, []int{1}},
		{nil, nil},
	}
	for _, c := range cases {
		got := toSlice(ReverseList(fromSlice(c.input)))
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("ReverseList(%v) = %v, want %v", c.input, got, c.want)
		}
	}
}
