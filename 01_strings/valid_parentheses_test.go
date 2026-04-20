package strings_practice

import "testing"

func TestIsValid(t *testing.T) {
	cases := []struct {
		s    string
		want bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true},
		{"[", false},
	}
	for _, c := range cases {
		if got := IsValid(c.s); got != c.want {
			t.Errorf("IsValid(%q) = %v, want %v", c.s, got, c.want)
		}
	}
}
