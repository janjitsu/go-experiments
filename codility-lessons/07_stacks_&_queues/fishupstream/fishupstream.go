package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(fish int) {
	*s = append(*s, fish)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return 0
	}
	index := len(*s) - 1
	last := (*s)[index]
	*s = (*s)[:index]
	return last
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return 0
	}
	return (*s)[len(*s)-1]
}

func Solution(A []int, B []int) int {
	// write your code in Go 1.4
	//B[P] == 1 fish going up meeting with A[P+1]
	const UPSTREAM = 1
	const DOWNSTREAM = 0

	var fishToBattle Stack
	var fishAlive = len(A)

	//fmt.Println(A)
	//fmt.Println(B)
	for i := 0; i < len(A); i++ {
		if B[i] == UPSTREAM {
			//fmt.Printf(">%d} ", A[i])
			fishToBattle.Push(A[i])
		} else {
			for fishToBattle.Peek() < A[i] && !fishToBattle.IsEmpty() { //upstream lost battle
				//fmt.Printf("x ")
				fishToBattle.Pop()
				fishAlive -= 1
			}
			if fishToBattle.Peek() > A[i] { //downstream lost battle
				//fmt.Printf("x ")
				fishAlive -= 1

			}
			//fmt.Printf("{%d< ", A[i])
		}
	}

	//fmt.Println()
	//fmt.Println(fishAlive)
	//fmt.Println(fishToBattle)

	return fishAlive
}
