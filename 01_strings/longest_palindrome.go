package strings_practice

// LongestPalindrome returns the longest palindromic substring.
// Input: "babad"  →  "bab" (or "aba")
// Input: "cbbd"   →  "bb"
func LongestPalindrome(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		odd := expand(s, i, i)
		even := expand(s, i, i+1)

		c := odd
		if len(even) > len(odd) {
			c = even
		}
		if len(c) > len(result) {
			result = c
		}
	}

	return result
}

func expand(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left --
		right ++
	}

	return s[left+1 : right]
}
