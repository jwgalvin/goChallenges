package longest_common_subsequence

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	cases := []struct {
		t1, t2 string
		want   int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
	}
	for _, c := range cases {
		if got := LongestCommonSubsequence(c.t1, c.t2); got != c.want {
			t.Errorf("LCS(%q, %q) = %d, want %d", c.t1, c.t2, got, c.want)
		}
	}
}
