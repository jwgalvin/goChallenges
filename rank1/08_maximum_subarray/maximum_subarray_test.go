package maximum_subarray

import "testing"

func TestMaxSubarray(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
		{[]int{-1, -2, -3}, -1},
	}
	for _, c := range cases {
		got := MaxSubarray(c.nums)
		if got != c.want {
			t.Errorf("MaxSubarray(%v) = %d, want %d", c.nums, got, c.want)
		}
	}
}
