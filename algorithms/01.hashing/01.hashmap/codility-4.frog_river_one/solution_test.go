package main

import (
	"alghoritms.com/check"
	"testing"
)

func TestSolution(t *testing.T) {
	check.CheckSolution(t, Solution, "test-input.txt", "test-output.txt")
}
