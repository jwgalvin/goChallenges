package rotate_array

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{1, 2}, 0, []int{1, 2}},
	}
	for _, c := range cases {
		nums := make([]int, len(c.nums))
		copy(nums, c.nums)
		Rotate(nums, c.k)
		if !reflect.DeepEqual(nums, c.want) {
			t.Errorf("Rotate(%v, %d) = %v, want %v", c.nums, c.k, nums, c.want)
		}
	}
}
