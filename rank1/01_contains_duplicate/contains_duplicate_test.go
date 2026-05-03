package contains_duplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"empty", []int{}, false},
		{"single element", []int{7}, false},
		{"all unique", []int{1, 2, 3, 4}, false},
		{"has duplicate", []int{1, 2, 3, 1}, true},
		{"all same", []int{5, 5, 5}, true},
		{"duplicate at end", []int{1, 2, 3, 4, 2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsDuplicate(tt.nums); got != tt.want {
				t.Errorf("ContainsDuplicate(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
