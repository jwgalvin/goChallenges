package letter_combinations

import "testing"

func TestLetterCombinations(t *testing.T) {
	got := LetterCombinations("23")
	if len(got) != 9 {
		t.Errorf("LetterCombinations(\"23\") returned %d results, want 9", len(got))
	}
	if LetterCombinations("") != nil {
		t.Error("LetterCombinations(\"\") should return nil")
	}
}
