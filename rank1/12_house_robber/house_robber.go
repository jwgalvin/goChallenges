package house_robber

// Rob returns the maximum amount you can rob from houses in a row
// without robbing two adjacent houses.
// Input: [2,7,9,3,1]  →  12  (2+9+1)
func Rob(nums []int) int {
	// TODO: implement (two-variable DP, O(n) time O(1) space)
	prev2 := 0
	prev1 := 0
	
	for _, num := range nums {
		current := max(prev1, num + prev2)
		prev2 = prev1
		prev1 = current
	}

	return prev1
}
