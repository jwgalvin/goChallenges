package merge_sorted

import (
	"slices"
	"testing"
)

func TestMergeSorted(t *testing.T) {
	tests := []struct {
		name string
		a, b []int
		want []int
	}{
		{"both empty", []int{}, []int{}, []int{}},
		{"a empty", []int{}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"b empty", []int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{"no overlap", []int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"interleaved", []int{1, 4, 7}, []int{2, 5, 8}, []int{1, 2, 4, 5, 7, 8}},
		{"duplicates across", []int{1, 2, 3}, []int{2, 3, 4}, []int{1, 2, 2, 3, 3, 4}},
		{"a longer", []int{1, 2, 3, 4, 5}, []int{3}, []int{1, 2, 3, 3, 4, 5}},
		{"b longer", []int{5}, []int{1, 2, 3, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeSorted(tt.a, tt.b)
			if !slices.Equal(got, tt.want) {
				t.Errorf("MergeSorted(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
