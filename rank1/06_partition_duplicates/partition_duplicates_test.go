package partition_duplicates

import (
	"reflect"
	"testing"
)

func TestPartitionDuplicates(t *testing.T) {
	tests := []struct {
		name       string
		nums       []int
		wantDups   map[int]int
		wantUnique map[int]int
	}{
		{
			name:       "empty",
			nums:       []int{},
			wantDups:   map[int]int{},
			wantUnique: map[int]int{},
		},
		{
			name:       "all unique",
			nums:       []int{1, 2, 3},
			wantDups:   map[int]int{},
			wantUnique: map[int]int{1: 1, 2: 1, 3: 1},
		},
		{
			name:       "all duplicates",
			nums:       []int{4, 4, 5, 5},
			wantDups:   map[int]int{4: 2, 5: 2},
			wantUnique: map[int]int{},
		},
		{
			name:       "mixed",
			nums:       []int{1, 2, 2, 3, 3, 3},
			wantDups:   map[int]int{2: 2, 3: 3},
			wantUnique: map[int]int{1: 1},
		},
		{
			name:       "single element",
			nums:       []int{7},
			wantDups:   map[int]int{},
			wantUnique: map[int]int{7: 1},
		},
		{
			name:       "triple duplicate mixed with unique",
			nums:       []int{9, 9, 9, 1, 2},
			wantDups:   map[int]int{9: 3},
			wantUnique: map[int]int{1: 1, 2: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDups, gotUnique := PartitionDuplicates(tt.nums)
			if !reflect.DeepEqual(gotDups, tt.wantDups) {
				t.Errorf("duplicates = %v, want %v", gotDups, tt.wantDups)
			}
			if !reflect.DeepEqual(gotUnique, tt.wantUnique) {
				t.Errorf("uniques = %v, want %v", gotUnique, tt.wantUnique)
			}
		})
	}
}
