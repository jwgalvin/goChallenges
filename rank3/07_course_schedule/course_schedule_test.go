package course_schedule

import "testing"

func TestCanFinish(t *testing.T) {
	cases := []struct {
		n    int
		pre  [][]int
		want bool
	}{
		{2, [][]int{{1, 0}}, true},
		{2, [][]int{{1, 0}, {0, 1}}, false},
		{1, nil, true},
	}
	for _, c := range cases {
		got := CanFinish(c.n, c.pre)
		if got != c.want {
			t.Errorf("CanFinish(%d, %v) = %v, want %v", c.n, c.pre, got, c.want)
		}
	}
}

func TestFindOrder(t *testing.T) {
	cases := []struct {
		n    int
		pre  [][]int
		want int // expected length of result (or -1 if impossible)
	}{
		{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, 4},
		{2, [][]int{{1, 0}}, 2},
		{2, [][]int{{1, 0}, {0, 1}}, -1}, // impossible, expect nil
		{1, nil, 1},
	}
	for _, c := range cases {
		got := FindOrder(c.n, c.pre)
		if c.want == -1 {
			if got != nil {
				t.Errorf("FindOrder(%d, %v) = %v, want nil (impossible)", c.n, c.pre, got)
			}
			continue
		}
		if len(got) != c.want {
			t.Errorf("FindOrder(%d, %v) returned %d courses, want %d", c.n, c.pre, len(got), c.want)
			continue
		}
		// Verify all prerequisites are satisfied
		pos := make(map[int]int, len(got))
		for i, course := range got {
			pos[course] = i
		}
		for _, pair := range c.pre {
			course, prereq := pair[0], pair[1]
			if pos[prereq] >= pos[course] {
				t.Errorf("FindOrder: prerequisite %d appears after course %d in order %v", prereq, course, got)
			}
		}
	}
}
