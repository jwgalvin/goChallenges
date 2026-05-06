package climbing_stairs

import "testing"

func TestClimbStairs(t *testing.T) {
	cases := []struct{ n, want int }{{2, 2}, {3, 3}, {4, 5}, {1, 1}}
	for _, c := range cases {
		if got := ClimbStairs(c.n); got != c.want {
			t.Errorf("ClimbStairs(%d) = %d, want %d", c.n, got, c.want)
		}
	}
}
