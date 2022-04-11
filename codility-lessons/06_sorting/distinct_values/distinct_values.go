package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	values := make(map[int]int)

	for i := 0; i < len(A); i++ {
		values[A[i]] += 1
	}

	return len(values)
}
