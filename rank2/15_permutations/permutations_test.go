package permutations

import "testing"

func TestPermute(t *testing.T) {
	got := Permute([]int{1, 2, 3})
	if len(got) != 6 {
		t.Errorf("Permute([1,2,3]) returned %d results, want 6", len(got))
	}
}
