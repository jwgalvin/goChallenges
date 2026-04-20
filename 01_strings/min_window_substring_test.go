package strings_practice

import "testing"

func TestMinWindow(t *testing.T) {
	cases := []struct {
		s, t string
		want string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "b", ""},
		{"aa", "aa", "aa"},
	}
	for _, c := range cases {
		got := MinWindow(c.s, c.t)
		if got != c.want {
			t.Errorf("MinWindow(%q, %q) = %q, want %q", c.s, c.t, got, c.want)
		}
	}
}
