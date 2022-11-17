package main

import (
	"fmt"
	"sort"
)

func SumTerms(A []int, target int) []int {
	//# first you sort O(n log n)
	//if the number is smaller than that value
	candidateX, candidateY, x, y := 0, 0, 0, 0
	sort.Ints(A)

	for i := 0; i < len(A); i++ {
		if A[i] < target {
			candidateX = A[i]
			candidateY = target - candidateX
		}
		if candidateY-A[i] == 0 {
			y = candidateY
			x = candidateX
		}
		fmt.Printf("A[i]=%d\nx=%d,y=%d,candidateX=%d,candidateY=%d\n", A[i], x, y, candidateX, candidateY)
	}

	return []int{x, y}
}

func main() {

}
