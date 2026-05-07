package sort_colors


// SortColors sorts an array containing only 0s, 1s, and 2s in-place
// (Dutch National Flag problem).
// Input: [2,0,2,1,1,0]  →  [0,0,1,1,2,2]
func SortColors(nums []int) {
	low, mid := 0, 0
	high := len(nums)-1

	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
	
	return

}
