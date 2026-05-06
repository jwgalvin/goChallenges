package edit_distance

import "testing"

func TestMinDistance(t *testing.T) {
	cases := []struct {
		w1, w2 string
		want   int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
		{"", "", 0},
	}
	for _, c := range cases {
		if got := MinDistance(c.w1, c.w2); got != c.want {
			t.Errorf("MinDistance(%q, %q) = %d, want %d", c.w1, c.w2, got, c.want)
		}
	}
}
