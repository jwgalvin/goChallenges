package arrays_practice

import "sort"

// Interval represents a closed interval [Start, End].
type Interval struct{ Start, End int }

// MergeIntervals merges all overlapping intervals.
// Input:  [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
func MergeIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	// TODO: implement merge logic
	return nil
}
