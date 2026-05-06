package longest_palindrome

import (
	"slices"
	"testing"
)

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
		if !slices.Contains(c.want, got) {
			t.Errorf("LongestPalindrome(%q) = %q, want one of %v", c.s, got, c.want)
		}
	}
}
