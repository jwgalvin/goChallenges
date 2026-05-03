package first_nonrepeating

import "testing"

func TestFirstNonRepeating(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"", -1},
		{"a", 0},
		{"aa", -1},
		{"ab", 0},
		{"aabb", -1},
		{"aabbc", 4},
		{"leetcode", 0},    // 'l' at 0 appears once
		{"loveleetcode", 2}, // 'v' at 2 appears once
		{"abcabc", -1},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := FirstNonRepeating(tt.input); got != tt.want {
				t.Errorf("FirstNonRepeating(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
