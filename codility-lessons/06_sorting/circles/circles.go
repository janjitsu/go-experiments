package solution

// you can also use imports, for example:
//import "fmt"
import "sort"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	circle_left := make([]int, len(A))
	circle_right := make([]int, len(A))
	circle_size := make([]int, len(A))

	for i := 0; i < len(A); i++ {
		circle_left[i] = i - A[i]
		circle_right[i] = i + A[i]
		circle_size[i] = A[i] * 2

	}
	sort.Ints(circle_left)
	sort.Ints(circle_right)

	/*
	   fmt.Println(circle_left)
	   fmt.Println(circle_right)
	   fmt.Println(circle_size)
	*/
	j := 0
	intersections := 0
	for i := 0; i < len(circle_left); i++ {
		for {
			if j < len(circle_right) && circle_right[i] >= circle_left[j] {
				intersections += j - i
				j++
			} else {
				break
			}
		}
		if intersections > 10000000 {
			return -1
		}

	}

	return intersections
}
