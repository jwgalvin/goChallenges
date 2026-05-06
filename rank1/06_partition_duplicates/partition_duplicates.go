package partition_duplicates

// PartitionDuplicates returns two maps from nums:
//   - duplicates: each value that appears more than once, mapped to its count
//   - uniques: each value that appears exactly once, mapped to its count (always 1)
func PartitionDuplicates(nums []int) (duplicates map[int]int, uniques map[int]int) {
	duplicates = make(map[int]int)
	uniques = make(map[int]int)

	for _, num := range nums {
		if _, ok := duplicates[num]; ok {
			duplicates[num]++
		} else if _, ok := uniques[num]; ok {
			duplicates[num] = 2
			delete(uniques, num)
		} else {
			uniques[num] = 1
		}
	}
	return
}
