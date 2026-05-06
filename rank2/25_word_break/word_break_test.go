package word_break

import "testing"

func TestWordBreak(t *testing.T) {
	cases := []struct {
		s    string
		dict []string
		want bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	}
	for _, c := range cases {
		if got := WordBreak(c.s, c.dict); got != c.want {
			t.Errorf("WordBreak(%q, %v) = %v, want %v", c.s, c.dict, got, c.want)
		}
	}
}
