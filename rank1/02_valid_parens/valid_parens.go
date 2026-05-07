package valid_parens

// IsValid returns true if every opening bracket in s has a matching closing
// bracket in the correct order. s contains only '(', ')', '{', '}', '[', ']'.
func IsValid(s string) bool {
	stack := []rune{}
	matching := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, c:= range s {
		if _, ok := matching[c]; ok {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if matching[top] != c {
				return false
			}
		}

	}
	return len(stack) == 0
}
