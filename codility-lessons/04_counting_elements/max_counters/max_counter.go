package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int, A []int) []int {
	// write your code in
	counters := make([]int, N)
	last_max := 0
	current_max := 0
	for i := 0; i < len(A); i++ {
		if A[i] > N {
			last_max = current_max
		} else {
			curr := A[i] - 1
			if counters[curr] < last_max {
				counters[curr] = last_max + 1
			} else {
				counters[curr] += 1
			}
			if current_max < counters[curr] {
				current_max = counters[curr]
			}
		}
		//fmt.Println(A[i])
		//fmt.Println(counters)
	}

	for j := 0; j < N; j++ {
		if counters[j] < last_max {
			counters[j] = last_max
		}
	}
	return counters
}
