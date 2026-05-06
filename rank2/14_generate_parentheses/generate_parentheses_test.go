package generate_parentheses

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
