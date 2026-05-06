package contains_duplicate


// ContainsDuplicate returns true if any value in nums appears more than once.
func ContainsDuplicate(nums []int) bool {
	seen := []int{}

	for _, num := range nums {
		if seen == nil {
			seen = append(seen, num)
		}
		for _, n := range seen {
			if n == num {
				return true
			}
		}
		seen = append(seen, num)
	}
	return false
}
