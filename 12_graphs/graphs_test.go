package graphs_practice

import "testing"

func TestNumIslands(t *testing.T) {
	cases := []struct {
		grid [][]byte
		want int
	}{
		{
			[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			1,
		},
		{
			[][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			3,
		},
		{[][]byte{}, 0},
	}
	for _, c := range cases {
		got := NumIslands(c.grid)
		if got != c.want {
			t.Errorf("NumIslands got %d, want %d", got, c.want)
		}
	}
}

func TestCanFinish(t *testing.T) {
	cases := []struct {
		n     int
		pre   [][]int
		want  bool
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
