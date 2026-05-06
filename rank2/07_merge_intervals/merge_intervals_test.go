package merge_intervals

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	cases := []struct {
		input []Interval
		want  []Interval
	}{
		{
			[]Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			[]Interval{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			[]Interval{{1, 4}, {4, 5}},
			[]Interval{{1, 5}},
		},
		{
			[]Interval{{1, 4}},
			[]Interval{{1, 4}},
		},
	}
	for _, c := range cases {
		got := MergeIntervals(c.input)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("MergeIntervals(%v) = %v, want %v", c.input, got, c.want)
		}
	}
}
