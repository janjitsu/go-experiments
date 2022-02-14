package main

import (
	"fmt"
	"os"
)

/* this function has a complexity of O(log n) */
func CheckUniqueCharacters(word string) bool {
	//explode characters
	char_count := make(map[string]int)

	for _, c := range word {
		char := string(c)
		if _, ok := char_count[char]; ok {
			return false
		}
		char_count[char] = 1
	}

	return true
}

func IsUnique(word string) {
	if CheckUniqueCharacters(word) {
		fmt.Println(fmt.Sprintf("[OK]    string \"%s\" has only unique characters", word))
	} else {
		fmt.Println(fmt.Sprintf("[ERROR] string \"%s\" has repeated characters", word))
	}
}

func main() {
	IsUnique(os.Args[1])
}
