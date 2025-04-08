package main

/*
 * https://app.codility.com/programmers/lessons/4-counting_elements/frog_river_one/
 */

func Solution(X int, A []int) int {
	//frog is in 0
	//X + 1 = where the frog want's to go ex: X=5, frog wants to go from 0 to 6
	//A positions where leaves are falling A[k] k = seconds.
	//  A[0] = 1
	//  A[1] = 3
	//  A[2] = 1
	//  A[3] = 4
	//  A[4] = 2
	//  A[5] = 3
	//  A[6] = 5
	//  A[7] = 4
	// in this example, given X=5, we need to know,
	//when is the first moment, that there are leaves enough from 1, to 5,
	//that is in position 6

	// frog can never jump
	return -1
}
