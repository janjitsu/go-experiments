package solution

// you can also use imports, for example:
import "fmt"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	prefix_sum := make([]int, len(A)+1)
	var avg_with_prev float64
	var avg_of_two float64
	min_avg_pos := 0
	left_pos := 0
	curr_avg := float64(A[0]+A[1]) / 2.0
	min_avg := curr_avg

	prefix_sum[0] = A[0]
	for i := 1; i < len(A); i++ {
		prefix_sum[i] = prefix_sum[i-1] + A[i]
	}

	fmt.Println(prefix_sum)

	for i := 2; i < len(A); i++ {
		avg_with_prev = (float64((prefix_sum[i+1] - prefix_sum[left_pos])) / float64(i-left_pos+1))

		avg_of_two = float64(A[i-1]+A[i]) / 2.0

		if avg_of_two < avg_with_prev {
			curr_avg = avg_of_two
			left_pos = i - 1
		} else {
			curr_avg = avg_with_prev
		}

		if curr_avg < min_avg {
			min_avg = curr_avg
			min_avg_pos = left_pos
		}

		fmt.Printf("[%d]avg_with_prev %2.2f = ((%d - %d) / (%d - %d + 1)))\n", i, avg_with_prev, prefix_sum[i], prefix_sum[left_pos], i, left_pos)
		fmt.Printf("[%d]avg_of_two %2.2f = (%d - %d) / 2.0 \n", i, avg_of_two, A[i-1], A[i])
		fmt.Printf("[%d]min_avg: %2.2f, min_avg_pos: %d\n", i, min_avg, min_avg_pos)
		fmt.Printf("-----------------\n")
	}

	return min_avg_pos
}
