package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	passing_east := make([]int, len(A))
	passing_cars := 0

	if A[0] == 0 {
		passing_east[0] = 1
	}

	for i := 1; i < len(A); i++ {
		if A[i] == 0 {
			passing_east[i] = 1 + passing_east[i-1]
		} else {
			passing_cars += passing_east[i-1]
			passing_east[i] = passing_east[i-1]
		}
		if passing_cars > 1000000000 {
			return -1
		}
	}

	return passing_cars
}
