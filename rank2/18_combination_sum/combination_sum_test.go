package combination_sum

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

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
	var b strings.Builder
	for _, v := range c {
		b.WriteRune(rune('0' + v))
		b.WriteByte(',')
	}
	return b.String()
}
