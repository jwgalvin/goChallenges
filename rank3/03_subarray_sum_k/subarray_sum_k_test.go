package subarray_sum_k

import "testing"

func TestSubarraySum(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 2, 3}, 3, 2},
		{[]int{1}, 0, 0},
		{[]int{-1, -1, 1}, 0, 1},
	}
	for _, c := range cases {
		got := SubarraySum(c.nums, c.k)
		if got != c.want {
			t.Errorf("SubarraySum(%v, %d) = %d, want %d", c.nums, c.k, got, c.want)
		}
	}
}
