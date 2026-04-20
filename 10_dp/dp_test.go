package dp_practice

import "testing"

func TestClimbStairs(t *testing.T) {
	cases := []struct{ n, want int }{{2, 2}, {3, 3}, {4, 5}, {1, 1}}
	for _, c := range cases {
		if got := ClimbStairs(c.n); got != c.want {
			t.Errorf("ClimbStairs(%d) = %d, want %d", c.n, got, c.want)
		}
	}
}

func TestRob(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{2, 1}, 2},
	}
	for _, c := range cases {
		if got := Rob(c.nums); got != c.want {
			t.Errorf("Rob(%v) = %d, want %d", c.nums, got, c.want)
		}
	}
}

func TestCoinChange(t *testing.T) {
	cases := []struct {
		coins  []int
		amount int
		want   int
	}{
		{[]int{1, 5, 11}, 15, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
	}
	for _, c := range cases {
		if got := CoinChange(c.coins, c.amount); got != c.want {
			t.Errorf("CoinChange(%v, %d) = %d, want %d", c.coins, c.amount, got, c.want)
		}
	}
}

func TestWordBreak(t *testing.T) {
	cases := []struct {
		s    string
		dict []string
		want bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	}
	for _, c := range cases {
		if got := WordBreak(c.s, c.dict); got != c.want {
			t.Errorf("WordBreak(%q, %v) = %v, want %v", c.s, c.dict, got, c.want)
		}
	}
}

func TestUniquePaths(t *testing.T) {
	cases := []struct{ m, n, want int }{{3, 7, 28}, {3, 2, 3}, {1, 1, 1}}
	for _, c := range cases {
		if got := UniquePaths(c.m, c.n); got != c.want {
			t.Errorf("UniquePaths(%d,%d) = %d, want %d", c.m, c.n, got, c.want)
		}
	}
}

func TestMinDistance(t *testing.T) {
	cases := []struct {
		w1, w2 string
		want   int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
		{"", "", 0},
	}
	for _, c := range cases {
		if got := MinDistance(c.w1, c.w2); got != c.want {
			t.Errorf("MinDistance(%q, %q) = %d, want %d", c.w1, c.w2, got, c.want)
		}
	}
}

func TestLengthOfLIS(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7}, 1},
	}
	for _, c := range cases {
		if got := LengthOfLIS(c.nums); got != c.want {
			t.Errorf("LengthOfLIS(%v) = %d, want %d", c.nums, got, c.want)
		}
	}
}

func TestLongestCommonSubsequence(t *testing.T) {
	cases := []struct {
		t1, t2 string
		want   int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
	}
	for _, c := range cases {
		if got := LongestCommonSubsequence(c.t1, c.t2); got != c.want {
			t.Errorf("LCS(%q, %q) = %d, want %d", c.t1, c.t2, got, c.want)
		}
	}
}

func TestMaxProduct(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{-2, 3, -4}, 24},
	}
	for _, c := range cases {
		if got := MaxProduct(c.nums); got != c.want {
			t.Errorf("MaxProduct(%v) = %d, want %d", c.nums, got, c.want)
		}
	}
}

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
