package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {

	var elementRegex = regexp.MustCompile(`"(?:[^"\\]|\\.)*"|\[.*?\]|[^,\[\]"]+`)

	file, err := os.Open("test-input.txt")
	if err != nil {
		log.Fatal("Error opening:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var parsed []interface{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		line = strings.TrimPrefix(line, "(")
		line = strings.TrimSuffix(line, ")")

		elements := elementRegex.FindAllString(line, -1)

		for _, el := range elements {
			val := parseValue(strings.TrimSpace(el))
			parsed = append(parsed, val)
		}
	}
	fmt.Printf("Parsed: %#v\n", parsed)

	result := callFunction(Solution, parsed)
	fmt.Printf("Input: %+v → Result: %v\n", parsed, result)

	/*
		for _, assertion := range testCases {
			//fmt.Printf("%+v\n", assertion)
			if (assertion.val) != assertion.expected {
				t.Fatalf("Failed asserting that IsUnique for %+v", assertion)
			}
		}

	*/
}
