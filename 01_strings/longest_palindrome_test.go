package strings_practice

import "testing"

func TestLongestPalindrome(t *testing.T) {
	cases := []struct {
		s    string
		want []string // acceptable answers
	}{
		{"babad", []string{"bab", "aba"}},
		{"cbbd", []string{"bb"}},
		{"a", []string{"a"}},
		{"racecar", []string{"racecar"}},
	}
	for _, c := range cases {
		got := LongestPalindrome(c.s)
		ok := false
		for _, w := range c.want {
			if got == w {
				ok = true
				break
			}
		}
		if !ok {
			t.Errorf("LongestPalindrome(%q) = %q, want one of %v", c.s, got, c.want)
		}
	}
}
