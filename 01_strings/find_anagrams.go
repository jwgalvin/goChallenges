package strings_practice

// FindAnagrams returns all start indices where p's anagrams appear in s.
// Input: s="cbaebabacd", p="abc"  →  [0, 6]

//s will be a longer string that could potentially have multiple anagrams of P.  
func FindAnagrams(s, p string) []int {
	// This solution assumes all lower case letters and no unicode, otherwise we would need to map runes
	// It also assumes that P will always have less characters, or the same number as S
	var pCount, sCount [26]int
	var result []int
	// This will avoid a panic if p has more chars than s
	if len(p) > len(s) {
    	return nil
	}
	// checks the first len(p) letters and sets their counts
	for i, ch := range p {
		pCount[ch-'a']++
		sCount[s[i]-'a']++
	}
	// if the counts match we append
	if pCount == sCount {
		result = append(result, 0)
	}

	// this loops through S given the window of len(p), window being the consecutive letters in the slice of the string
	for i := len(p); i < len(s); i++ {
		// this takes the character, figures out how far it is from 'a', and use that distance as the array index.
		sCount[s[i]-'a'] ++
		sCount[s[i - len(p)]-'a']--
		// if its an anagram add to the result
		if pCount == sCount {
			result = append(result, i-len(p)+1)
		}
	}

	return result
}
