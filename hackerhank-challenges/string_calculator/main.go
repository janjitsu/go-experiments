package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// "5 + 7 - 2 + 10 * 6 / 5" = 22
// input will be a string with the formula
// output will be calculated result
// "1 + 1"

func stringsToFloats(strings []string) []float64 {
	floats := make([]float64, len(strings))
	for i := 0; i < len(strings); i++ {
		number, _ := strconv.Atoi(strings[i])
		floats[i] = float64(number)
	}

	return floats
}

func calculateSumsAndSubtractions(text string) float64 {

	var result float64

	findOperators, _ := regexp.Compile("\\+|\\-")
	findNumbers, _ := regexp.Compile("[0-9]+")

	operators := findOperators.FindAllString(text, -1)
	numbers := stringsToFloats(findNumbers.FindAllString(text, -1))

	result = numbers[0]
	for i := 0; i < len(operators); i++ {
		if operators[i] == "+" {
			result = result + numbers[i+1]
		}
		if operators[i] == "-" {
			result = result - numbers[i+1]
		}
	}

	return result
}

func replaceMultiplications(text string) string {
	multiplicatons, _ := regexp.Compile("([0-9]+) \\* ([0-9]+)")
	operations := multiplicatons.FindAllString(text, -1)
	findNumbers, _ := regexp.Compile("[0-9]+")

	for i := 0; i < len(operations); i++ {
		fmt.Println(operations[i])
		numbers := stringsToFloats(findNumbers.FindAllString(operations[i], -1))
		result := numbers[0] * numbers[1]
		text = strings.Replace(text, operations[i], fmt.Sprintf("%2.0f", result), -1)
	}

	return text
}

func replaceDivisions(text string) string {
	divisions, _ := regexp.Compile("([0-9]+) \\/ ([0-9]+)")
	operations := divisions.FindAllString(text, -1)
	findNumbers, _ := regexp.Compile("[0-9]+")

	for i := 0; i < len(operations); i++ {
		fmt.Println(operations[i])
		numbers := stringsToFloats(findNumbers.FindAllString(operations[i], -1))
		result := numbers[0] / numbers[1]
		text = strings.Replace(text, operations[i], fmt.Sprintf("%2.0f", result), -1)
	}

	return text
}

func main() {
	//exp := "5 + 7 - 2 + 10 * 6 / 5"
	exp := "5 * 7 - 2 + 10 * 6 / 5"
	exp = replaceMultiplications(exp)
	exp = replaceDivisions(exp)

	result := calculateSumsAndSubtractions(exp)

	fmt.Println("Result is ", result)
}

// how to approach practical live code problems?
// try to see all the hidden rules for the solution to work
