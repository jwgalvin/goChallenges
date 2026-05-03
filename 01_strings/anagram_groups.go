package strings_practice

import "sort"

// GroupAnagrams groups strings that are anagrams of each other.
// Input:  ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
func GroupAnagrams(strs []string) [][]string {
	grouping := make(map[string][]string)
	result := [][]string{}
	for _, word := range strs {
		// convert to bytes
		b := []byte(word)
		// sort the bytes by ASC order
		sort.Slice(b, func(i,j int) bool { return b[i] < b[j] })
		// create the key so it is the sorted byte value of the word
		key :=string(b)
		// appends the word to the map with the key of its sorted slice
		grouping[key] = append(grouping[key], word) 
	}

	// iterates through each group in grouping and adds the group to the result, effectively aggregating words that are anagrams
	for _, group := range grouping{
		result = append(result, group)
	}

	return result
}
