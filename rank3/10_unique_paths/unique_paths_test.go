package unique_paths

import "testing"

func TestUniquePaths(t *testing.T) {
	cases := []struct{ m, n, want int }{{3, 7, 28}, {3, 2, 3}, {1, 1, 1}}
	for _, c := range cases {
		if got := UniquePaths(c.m, c.n); got != c.want {
			t.Errorf("UniquePaths(%d,%d) = %d, want %d", c.m, c.n, got, c.want)
		}
	}
}
