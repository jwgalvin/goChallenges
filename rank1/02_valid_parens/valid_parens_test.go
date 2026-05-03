package valid_parens

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", true},
		{"()", true},
		{"()[]{}", true},
		{"{[]}", true},
		{"((()))", true},
		{"(]", false},
		{"([)]", false},
		{"{", false},
		{"}", false},
		{"((())", false},
		{"]", false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := IsValid(tt.input); got != tt.want {
				t.Errorf("IsValid(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
