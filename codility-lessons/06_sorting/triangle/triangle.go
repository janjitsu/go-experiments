package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	sort.Ints(A)

	for i := 0; i < len(A)-2; i++ {
		if (A[i]+A[i+1] > A[i+2]) &&
			(A[i+1]+A[i+2] > A[i]) &&
			(A[i+2]+A[i] > A[i+1]) {
			return 1
		}

	}

	return 0
}
