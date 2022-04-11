package solution

// you can also use imports, for example:
//import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	pairs := make(map[int]int)
	for i := 0; i < len(A); i++ {
		num := A[i]
		if _, ok := pairs[num]; ok {
			pairs[num] += 1
		} else {
			pairs[num] = 1
		}
	}
	//fmt.Println(pairs)
	for k, v := range pairs {
		if v%2 == 1 {
			return k
		}
	}

	return 1
}
