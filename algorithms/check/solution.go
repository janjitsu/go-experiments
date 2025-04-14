package check

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func callFunction(fn interface{}, args []interface{}) interface{} {
	fnValue := reflect.ValueOf(fn)

	var reflectArgs []reflect.Value
	for _, arg := range args {
		reflectArgs = append(reflectArgs, reflect.ValueOf(arg))
	}
	results := fnValue.Call(reflectArgs)

	return results[0].Interface()
}

func parseValue(s string) interface{} {
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		return strings.Trim(s, "\"")
	}

	// Boolean
	if s == "true" || s == "false" {
		return s == "true"
	}

	// Array
	if strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") {
		s = strings.Trim(s, "[]")
		parts := strings.Split(s, ",")
		var nums []int
		for _, part := range parts {
			n, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				return s // fallback to raw string if parse fails
			}
			nums = append(nums, n)
		}
		return nums
	}

	// Int
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}

	// Float
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	}

	// Default: string
	return s
}

func parseValuesFile(inputFile string) [][]interface{} {
	var elementRegex = regexp.MustCompile(`"(?:[^"\\]|\\.)*"|\[.*?\]|[^,\[\]"]+`)

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal("Error opening:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var inputs [][]interface{}
	for scanner.Scan() {
		var parsed []interface{}
		line := strings.TrimSpace(scanner.Text())
		line = strings.TrimPrefix(line, "(")
		line = strings.TrimSuffix(line, ")")

		elements := elementRegex.FindAllString(line, -1)

		for _, el := range elements {
			val := parseValue(strings.TrimSpace(el))
			parsed = append(parsed, val)
		}
		inputs = append(inputs, parsed)
	}
	return inputs
}

func CheckSolution(t *testing.T, fn interface{}, inputFile string,
	outputFile string) {
	inputs := parseValuesFile(inputFile)
	outputs := parseValuesFile(outputFile)
	for i, input := range inputs {
		result := callFunction(fn, input)
		expected := outputs[i][0]
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	}
}
