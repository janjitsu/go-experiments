package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {

	ocurrences := make([]int, len(A)+1)

	//fmt.Println(ocurrences)

	for i := 0; i < len(ocurrences); i++ {
		ocurrences[i] = i + 1
	}

	for j := 0; j < len(A); j++ {
		ocurrences[A[j]-1] = 0
	}

	for h := 0; h < len(ocurrences); h++ {
		if ocurrences[h] != 0 {
			return h + 1
		}
	}

	return 1
}
