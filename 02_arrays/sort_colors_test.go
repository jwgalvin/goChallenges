package arrays_practice

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	cases := []struct {
		nums []int
		want []int
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{2, 0, 1}, []int{0, 1, 2}},
		{[]int{0}, []int{0}},
		{[]int{1, 2, 0}, []int{0, 1, 2}},
	}
	for _, c := range cases {
		nums := make([]int, len(c.nums))
		copy(nums, c.nums)
		SortColors(nums)
		if !reflect.DeepEqual(nums, c.want) {
			t.Errorf("SortColors(%v) = %v, want %v", c.nums, nums, c.want)
		}
	}
}
