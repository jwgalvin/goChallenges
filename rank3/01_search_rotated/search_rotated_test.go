package search_rotated

import "testing"

func TestSearchRotated(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{1}, 1, 0},
	}
	for _, c := range cases {
		got := SearchRotated(c.nums, c.target)
		if got != c.want {
			t.Errorf("SearchRotated(%v, %d) = %d, want %d", c.nums, c.target, got, c.want)
		}
	}
}
