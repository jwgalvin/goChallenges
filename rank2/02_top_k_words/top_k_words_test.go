package top_k_words

import (
	"slices"
	"testing"
)

func TestTopKWords_Basic(t *testing.T) {
	words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
	got := TopKWords(words, 2)
	want := []string{"apple", "banana"}
	if !slices.Equal(got, want) {
		t.Errorf("TopKWords(..., 2) = %v, want %v", got, want)
	}
}

func TestTopKWords_AlphabeticalTiebreak(t *testing.T) {
	// "a" and "b" both appear twice; "c" once — top 2 should be ["a","b"]
	words := []string{"b", "a", "b", "a", "c"}
	got := TopKWords(words, 2)
	want := []string{"a", "b"}
	if !slices.Equal(got, want) {
		t.Errorf("TopKWords() = %v, want %v", got, want)
	}
}

func TestTopKWords_KExceedsUniqueCount(t *testing.T) {
	words := []string{"x", "y", "x"}
	got := TopKWords(words, 10)
	if len(got) != 2 {
		t.Errorf("got %d words, want 2 (x and y)", len(got))
	}
}

func TestTopKWords_Empty(t *testing.T) {
	got := TopKWords([]string{}, 3)
	if len(got) != 0 {
		t.Errorf("expected empty result, got %v", got)
	}
}

func TestTopKWords_KIsZero(t *testing.T) {
	got := TopKWords([]string{"a", "b"}, 0)
	if len(got) != 0 {
		t.Errorf("k=0: expected empty result, got %v", got)
	}
}

func TestTopKWords_SingleWord(t *testing.T) {
	got := TopKWords([]string{"go", "go", "go"}, 1)
	want := []string{"go"}
	if !slices.Equal(got, want) {
		t.Errorf("TopKWords() = %v, want %v", got, want)
	}
}
