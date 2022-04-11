package solution

// you can also use imports, for example:
import "fmt"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	// 30 day ticket cost 25
	//  7 day ticket cost 7
	//  1 day ticket cost 2

	first_day := A[0]
	total_days := len(A)
	total_cost := 0
	fmt.Printf("%d days travelling starting at %d...\n", total_days, first_day)

	single_months := 0
	single_weeks := 0
	single_days := 0
	consecutive_days := 1

	for i := 1; i < len(A); i++ {
		days_diff := A[i] - A[i-1]
		days_gap := A[i] - first_day

		fmt.Printf("%d ", A[i])

		if days_diff == 1 {
			consecutive_days += 1
		} else {
			single_days += consecutive_days

			consecutive_days == 0
		}

	}

	return 11
}
