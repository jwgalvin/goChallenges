package partition_equal_subset

import "testing"

func TestCanPartition(t *testing.T) {
	cases := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 5, 11, 5}, true},
		{[]int{1, 2, 3, 5}, false},
	}
	for _, c := range cases {
		if got := CanPartition(c.nums); got != c.want {
			t.Errorf("CanPartition(%v) = %v, want %v", c.nums, got, c.want)
		}
	}
}
