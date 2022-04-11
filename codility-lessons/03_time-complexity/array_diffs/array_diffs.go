package solution

// you can also use imports, for example:
// import "fmt"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func AbsInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func Solution(A []int) int {
	// write your code in Go 1.4
	arr_size := len(A)

	left_sum := make([]int, arr_size)
	right_sum := make([]int, arr_size)

	left_sum[0] = 0
	for i := 1; i < len(A); i++ {
		left_sum[i] = left_sum[i-1] + A[i-1]
	}

	right_sum[arr_size-1] = A[arr_size-1]
	for j := len(A) - 2; j >= 0; j-- {
		right_sum[j] = right_sum[j+1] + A[j]
	}

	all_diffs := make([]int, arr_size)
	min_dif := int(^uint(0) >> 1) //max integer

	for k := 1; k < arr_size; k++ {
		all_diffs[k] = AbsInt(left_sum[k] - right_sum[k])
		if min_dif > all_diffs[k] {
			min_dif = all_diffs[k]
		}
	}

	/*
		fmt.Println(left_sum)
		fmt.Println(right_sum)
		fmt.Println(all_diffs)
		fmt.Println(min_dif)
	*/
	return min_dif
}
