package strings_practice

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	cases := []struct {
		s, p string
		want []int
	}{
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},
		{"aa", "bb", nil},
	}
	for _, c := range cases {
		got := FindAnagrams(c.s, c.p)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("FindAnagrams(%q, %q) = %v, want %v", c.s, c.p, got, c.want)
		}
	}
}
