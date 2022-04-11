package solution

// you can also use imports, for example:
//import "fmt"
//import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	}
	index := len(*s) - 1
	last := (*s)[index]
	*s = (*s)[:index]
	return last
}

func (s *Stack) Peek() string {
	if s.IsEmpty() {
		return ""
	}
	return (*s)[len(*s)-1]
}

func Solution(S string) int {
	// write your code in Go 1.4

	var stackOpen Stack

	//var stackClosed Stack
	oposing := map[string]string{
		"}": "{", "]": "[", ")": "(",
	}

	for _, c := range S {
		char := string(c)
		if char == "}" || char == "]" || char == ")" {
			if stackOpen.IsEmpty() {
				return 0
			}
			if stackOpen.Peek() == oposing[char] {
				stackOpen.Pop()
			}
		} else {
			stackOpen.Push(char)
		}
	}
	if stackOpen.IsEmpty() {
		return 1
	}
	return 0
}
