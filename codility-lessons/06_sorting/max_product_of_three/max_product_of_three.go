package solution

// you can also use imports, for example:
//import "fmt"
import (
	"math"
	"sort"
)

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	sort.Ints(A)
	//fmt.Println(A)
	size := len(A)

	if size == 3 {
		return A[0] * A[1] * A[2]
	}

	last1 := A[size-1]
	last2 := A[size-2]
	last3 := A[size-3]
	first1 := A[0]
	first2 := A[1]
	max1 := float64(last1 * last2 * last3)
	max2 := float64(last1 * first1 * first2)

	return int(math.Max(max1, max2))
}
