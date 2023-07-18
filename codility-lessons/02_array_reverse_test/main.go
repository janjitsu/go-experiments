package main

import "fmt"

func reverse(a []int, k int) []int {
	for i := 0; i < k; i++ {
		a = append([]int{a[len(a)-1]}, a[:len(a)-1]...)
	}

	return a
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(a)
	fmt.Println(a[:len(a)-1])
	fmt.Println(a[:0])
	fmt.Println(a[0 : len(a)-2])
	fmt.Println(reverse(a, 1))
}
