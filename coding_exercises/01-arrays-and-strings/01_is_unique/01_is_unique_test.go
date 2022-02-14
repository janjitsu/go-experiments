package main

import (
	"testing"
)

type StringTest struct {
	val      string
	expected bool
}

func TestCheckUniqueCharacters(t *testing.T) {

	strings_to_test := []StringTest{
		{"teucu", false},
		{"abcde123", true},
		{"rsckuarsntuafnt", false},
		{"12934arskzstetmvndg20p9uupetiafepmy234ul95upuarmteitnqwfpqn4n", false},
		{"1234567890qwfpgjluy;arstdhneiozxcvbkm,.", true},
	}

	for _, assertion := range strings_to_test {
		//fmt.Printf("%+v\n", assertion)
		if CheckUniqueCharacters(assertion.val) != assertion.expected {
			t.Fatalf("Failed asserting that IsUnique for %+v", assertion)
		}
	}
}
