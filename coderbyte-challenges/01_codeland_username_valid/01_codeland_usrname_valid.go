package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func CodelandUsernameValidation(str string) string {

	r, _ := regexp.Compile("[a-zA-Z][a-zA-Z0-9_]{2,23}[a-zA-Z0-9]")
	if r.MatchString(str) {
		fmt.Println("did a match!")
	}

	return strconv.FormatBool(r.MatchString(str))
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(CodelandUsernameValidation("art_"))
	fmt.Println(CodelandUsernameValidation("art_1"))
}
