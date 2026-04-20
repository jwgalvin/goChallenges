package backtracking_practice

import (
	"reflect"
	"sort"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	got := GenerateParenthesis(3)
	want := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	sort.Strings(got)
	sort.Strings(want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GenerateParenthesis(3) = %v, want %v", got, want)
	}
}

func TestPermute(t *testing.T) {
	got := Permute([]int{1, 2, 3})
	if len(got) != 6 {
		t.Errorf("Permute([1,2,3]) returned %d results, want 6", len(got))
	}
}

func TestSubsets(t *testing.T) {
	got := Subsets([]int{1, 2, 3})
	if len(got) != 8 {
		t.Errorf("Subsets([1,2,3]) returned %d results, want 8", len(got))
	}
}

func TestCombinationSum(t *testing.T) {
	got := CombinationSum([]int{2, 3, 6, 7}, 7)
	// Normalize: sort each combo and sort the list
	for _, c := range got {
		sort.Ints(c)
	}
	sort.Slice(got, func(i, j int) bool {
		return fmt_combo(got[i]) < fmt_combo(got[j])
	})
	want := [][]int{{2, 2, 3}, {7}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CombinationSum got %v, want %v", got, want)
	}
}

func fmt_combo(c []int) string {
	s := ""
	for _, v := range c {
		s += string(rune('0'+v)) + ","
	}
	return s
}

func TestLetterCombinations(t *testing.T) {
	got := LetterCombinations("23")
	if len(got) != 9 {
		t.Errorf("LetterCombinations(\"23\") returned %d results, want 9", len(got))
	}
	if LetterCombinations("") != nil {
		t.Error("LetterCombinations(\"\") should return nil")
	}
}
