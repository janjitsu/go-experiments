package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, K int) []int {
	// write your code in Go 1.4
	r := K

	if r > len(A) {
		r = (K % len(A))
	}

	if r == len(A) || r <= 0 {
		return A
	}

	return append(A[r-1:], A[:r-1]...)
}


package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
func Solution(A []int, K int) []int {
    // write your code in Go 1.4
    //fmt.Println([]int{A[len(A)-1]})
    //fmt.Println(A[:len(A)-1])
    if (len(A) == 0 || K == 0) {
        return A
    }

    for i := 0; i < K; i++ {
        A = append([]int{A[len(A)-1]},A[:len(A)-1]...)
    }

    return A
}


