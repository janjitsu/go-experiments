package solution

// you can also use imports, for example:
import (
	"fmt"
	"sort"
)

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func find_leader(A []int, candidate int) int {

	leader, value, freq := -1, 0, 0
	sorted := append([]int{}, A...)
	sort.Ints(sorted)

	if candidate == -1 { //find candidate
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
	}

	for j := 0; j < len(sorted); j++ {
		if sorted[j] == candidate {
			freq += 1
		}
	}

	if freq > (len(sorted) / 2) {
		leader = candidate
	}

	return leader
}

func Solution(A []int) int {
	// write your code in Go 1.4
	// first we get the leader
	size := len(A)
	leader_candidate := find_leader(A, -1)
	leader_sum_head := make([]int, size)
	leader_sum_tail := make([]int, size)

	equi_leaders, leader_count_head, leader_count_tail := 0, 0, 0

	if leader_candidate == -1 {
		return 0
	}

	for i := 0; i < size; i++ {
		if A[i] == leader_candidate {
			leader_count_head += 1
		}
		if A[size-1-i] == leader_candidate {
			leader_count_tail += 1
		}
		leader_sum_head[i] = leader_count_head
		leader_sum_tail[i] = leader_count_tail
	}

	for j := 0; j < size; j++ {

		leader_head := leader_sum_head[j] > ((1 + j) / 2)
		leader_tail := leader_sum_tail[size-1-j] > ((size - 1 - j) / 2)
		fmt.Printf("[head[%d]] %d > %d && ", j, leader_sum_head[j], (1+j)/2)
		fmt.Printf("[tail[%d]] %d > %d", j, leader_sum_tail[size-1-j], (size-1-j)/2)

		if leader_head && leader_tail {
			equi_leaders += 1
			fmt.Printf(" = yes")
		}
		fmt.Println()
	}

	fmt.Println(A)
	fmt.Println(leader_sum_head)
	fmt.Println(leader_sum_tail)

	return equi_leaders
}
