package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4

	sorted := append([]int{}, A...)
	sort.Ints(sorted)

	leader, candidate, freq, value := -1, -1, 0, 0

	for i := 0; i < len(sorted); i++ {
		if freq == 0 {
			freq += 1
			value = sorted[i]
		} else {
			if value == sorted[i] {
				freq += 1
			} else {
				freq -= 1
			}
		}
	}

	if freq > 0 {
		candidate = value
	}
	freq = 0
	for j := 0; j < len(sorted); j++ {
		if sorted[j] == candidate {
			freq += 1
		}
	}

	if freq > (len(sorted) / 2) {
		leader = candidate
	}
	for k := 0; k < len(A); k++ {
		if A[k] == leader {
			return k
		}
	}

	return -1
}
