package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4

	counts := make(map[int]int, len(A))
	high_num := 0

	// count ocurrences
	for i := 0; i < len(A); i++ {
		counts[A[i]] += 1
		if A[i] > high_num {
			high_num = A[i]
		}
	}

	// check if has sequence
	for i := 1; i <= high_num; i++ {
		if v := counts[i]; v != 1 {
			return 0
		}
	}

	return 1
}
