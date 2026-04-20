package arrays_practice

import (
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
	}
	for _, c := range cases {
		got := TopKFrequent(c.nums, c.k)
		sort.Ints(got)
		sort.Ints(c.want)
		ok := len(got) == len(c.want)
		if ok {
			for i := range got {
				if got[i] != c.want[i] {
					ok = false
					break
				}
			}
		}
		if !ok {
			t.Errorf("TopKFrequent(%v, %d) = %v, want %v", c.nums, c.k, got, c.want)
		}
	}
}
