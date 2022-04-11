package solution

// you can also use imports, for example:
//import "fmt"
import "strings"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string, P []int, Q []int) []int {
	// write your code in Go 1.4
	dna_stripe := strings.Split(S, "")
	prefix_sum_a := make([]int, len(dna_stripe)+1)
	prefix_sum_c := make([]int, len(dna_stripe)+1)
	prefix_sum_g := make([]int, len(dna_stripe)+1)
	min_factor := make([]int, len(P))

	for k := 0; k < len(dna_stripe); k++ {
		a, c, g := 0, 0, 0
		if dna_stripe[k] == "A" {
			a = 1
		}
		if dna_stripe[k] == "C" {
			c = 1
		}
		if dna_stripe[k] == "G" {
			g = 1
		}
		prefix_sum_a[k+1] = prefix_sum_a[k] + a
		prefix_sum_c[k+1] = prefix_sum_c[k] + c
		prefix_sum_g[k+1] = prefix_sum_g[k] + g
	}

	/*
	   fmt.Println(prefix_sum_a)
	   fmt.Println(prefix_sum_c)
	   fmt.Println(prefix_sum_g)
	*/

	for i := 0; i < len(P); i++ {
		if prefix_sum_a[Q[i]+1]-prefix_sum_a[P[i]] > 0 {
			min_factor[i] = 1
		} else if prefix_sum_c[Q[i]+1]-prefix_sum_c[P[i]] > 0 {
			min_factor[i] = 2
		} else if prefix_sum_g[Q[i]+1]-prefix_sum_g[P[i]] > 0 {
			min_factor[i] = 3
		} else {
			min_factor[i] = 4
		}
	}
	return min_factor
}
