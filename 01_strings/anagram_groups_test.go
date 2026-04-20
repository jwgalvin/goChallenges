package strings_practice

import (
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		input    []string
		wantSize int // number of groups
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, 3},
		{[]string{"a"}, 1},
		{[]string{""}, 1},
		{[]string{"", ""}, 1},
	}
	for _, tt := range tests {
		got := GroupAnagrams(tt.input)
		if len(got) != tt.wantSize {
			t.Errorf("GroupAnagrams(%v) = %d groups, want %d", tt.input, len(got), tt.wantSize)
		}
		// verify every word is accounted for
		total := 0
		for _, g := range got {
			total += len(g)
		}
		if total != len(tt.input) {
			t.Errorf("word count mismatch: got %d, want %d", total, len(tt.input))
		}
		_ = sort.Search // suppress unused import
	}
}
