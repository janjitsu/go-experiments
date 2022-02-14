package main

import (
	"fmt"
	"os"
)

/***********************************************************************************
* Given two strings, write a method to decide if one is a permutation of the other *
***********************************************************************************/

/* this function has a complexity of O(n)+O(log n) O(n log n)? */
func CheckIfPermutation(word string, permutation string) bool {

	word_chars := make(map[string]int)
	for _, c := range word {
		char := string(c)
		if _, ok := word_chars[char]; ok {
			word_chars[char] = word_chars[char] + 1
		} else {
			word_chars[char] = 1
		}
	}
	for _, c := range permutation {
		char := string(c)
		if count, ok := word_chars[char]; !ok {
			return false
		} else if count > 1 {
			word_chars[char] = word_chars[char] - 1
		} else {
			delete(word_chars, char)
		}
	}
	return true
}

func IsPermutation(word string, permutation string) {
	if CheckIfPermutation(word, permutation) {
		fmt.Printf("[ok]    %s is permutation of %s\n", permutation, word)
	} else {
		fmt.Printf("[error] %s is not permutation of %s\n", permutation, word)
	}
}

func main() {
	IsPermutation(os.Args[1], os.Args[2])
}
