package prefix_sum

import (
	"slices"
	"testing"
)

func TestPrefixSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"empty", []int{}, []int{}},
		{"single", []int{5}, []int{5}},
		{"all positive", []int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{"all negative", []int{-1, -2, -3}, []int{-1, -3, -6}},
		{"mixed", []int{5, -3, 2, -1}, []int{5, 2, 4, 3}},
		{"zeros", []int{0, 0, 0}, []int{0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PrefixSum(tt.nums)
			if !slices.Equal(got, tt.want) {
				t.Errorf("PrefixSum(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
