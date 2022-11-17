package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestSumTermsNegativeAndPositiveTerms(t *testing.T) {
	testArray := []int{-7, -3, -2, 1, 9, 15, 16, 17, 18}
	testInput := 9
	expected := []int{16, -7}
	testResult := SumTerms(testArray, testInput)

	sort.Ints(expected)
	sort.Ints(testResult)

	if !reflect.DeepEqual(testResult, expected) {
		t.Errorf(
			"Error testing with input SumTerms(%+v, %d) = %+v, expected %+v",
			testArray, testInput, testResult, expected,
		)
	}
}

func TestSumTermsSmallTerms(t *testing.T) {
	testArray := []int{-7, 16}
	testInput := 9
	expected := []int{16, -7}
	testResult := SumTerms(testArray, testInput)

	sort.Ints(expected)
	sort.Ints(testResult)

	if !reflect.DeepEqual(testResult, expected) {
		t.Errorf(
			"Error testing with input SumTerms(%+v, %d) = %+v, expected %+v",
			testArray, testInput, testResult, expected,
		)
	}
}

func TestSumTermsOnlyPositives(t *testing.T) {
	testArray := []int{0, 10, 15, 2, 8, 7, 5, 2}
	testInput := 4
	expected := []int{2, 2}
	testResult := SumTerms(testArray, testInput)

	sort.Ints(expected)
	sort.Ints(testResult)

	if !reflect.DeepEqual(testResult, expected) {
		t.Errorf(
			"Error testing with input SumTerms(%+v, %d) = %+v, expected %+v",
			testArray, testInput, testResult, expected,
		)
	}
}
