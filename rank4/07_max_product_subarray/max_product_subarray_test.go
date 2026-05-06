package max_product_subarray

import "testing"

func TestMaxProduct(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{-2, 3, -4}, 24},
	}
	for _, c := range cases {
		if got := MaxProduct(c.nums); got != c.want {
			t.Errorf("MaxProduct(%v) = %d, want %d", c.nums, got, c.want)
		}
	}
}
