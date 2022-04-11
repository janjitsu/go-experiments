package solution

// you can also use imports, for example:
//import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {

	count := make([]int, 1000001)
	max_num := 0

	for i := 0; i < len(A); i++ {
		if A[i] > 0 {
			count[A[i]] = 1
			if A[i] > max_num {
				max_num = A[i]
			}
		}
	}
	//fmt.Println(count)

	for j := 0; j <= max_num; j++ {
		if ok := count[j+1]; ok == 0 {
			return j + 1
		}
		//fmt.Printf("curr: %d\n",j)
	}

	return 1
}
