package main

import (
	"testing"
)

type Assertion struct {
	word        string
	permutation string
	expected    bool
}

func TestCheckIfPermutation(t *testing.T) {
	words_to_test := []Assertion{
		{"teucu", "cuteu", true},
		{"word", "drow", true},
		{"word", "drowwwwww", false},
	}

	for _, assertion := range words_to_test {
		//fmt.Printf("%+v\n", assertion)
		if CheckIfPermutation(assertion.word, assertion.permutation) != assertion.expected {
			t.Fatalf("Failed asserting that IsPermutation for %+v", assertion)
		}
	}
}
