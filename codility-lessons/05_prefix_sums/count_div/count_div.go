package solution

// you can also use imports, for example:
//import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A int, B int, K int) int {
	count_div := (B/K - A/K)
	tail_div := 0

	if A%K == 0 {
		tail_div += 1
	}

	return count_div + tail_div
}
