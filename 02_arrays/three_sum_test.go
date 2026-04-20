package arrays_practice

import (
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	normalize := func(res [][]int) [][]int {
		for _, r := range res {
			sort.Ints(r)
		}
		sort.Slice(res, func(i, j int) bool {
			for k := 0; k < len(res[i]) && k < len(res[j]); k++ {
				if res[i][k] != res[j][k] {
					return res[i][k] < res[j][k]
				}
			}
			return false
		})
		return res
	}

	cases := []struct {
		nums []int
		want [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, nil},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}
	for _, c := range cases {
		got := normalize(ThreeSum(c.nums))
		want := normalize(c.want)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("ThreeSum(%v) = %v, want %v", c.nums, got, want)
		}
	}
}
