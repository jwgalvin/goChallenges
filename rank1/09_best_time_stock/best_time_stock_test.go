package best_time_stock

import "testing"

func TestMaxProfit(t *testing.T) {
	cases := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2}, 1},
		{[]int{2, 4, 1}, 2},
	}
	for _, c := range cases {
		got := MaxProfit(c.prices)
		if got != c.want {
			t.Errorf("MaxProfit(%v) = %d, want %d", c.prices, got, c.want)
		}
	}
}
