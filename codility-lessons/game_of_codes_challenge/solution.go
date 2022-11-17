package solution

// you can also use imports, for example:
import "strings"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string) int {
	str := strings.Split(S, "")
	var blocks []string

	blocks = append(blocks, str[0])

	for i := 1; i < len(str)-1; i++ {
		if str[i] == blocks[len(blocks)-1] {
			blocks = append(blocks, str[i])

		} else if str[i+1] == str[i] {
			blocks = append(blocks, str[i])

		} else if i < len(str)-2 && str[i+2] == str[i] {
			blocks = append(blocks, str[i])

		}
	}
	return len(blocks) + 1
}
