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

func (s *Stack) Push(rectangle int) {
	*s = append(*s, rectangle)
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

func Solution(H []int) int {
	// write your code in Go 1.4

	//width := len(H)
	rectangles := 1
	var rectangles_bellow Stack
	rectangles_bellow.Push(H[0])

	/*
	   fmt.Printf("build is %d meters long\n",width)
	   for j := 0; j <= H[0]; j++ {
	       fmt.Printf("-")
	   }
	   fmt.Printf("%d|\n",H[0])
	*/
	for i := 1; i < len(H); i++ {
		/*
		   for j := 0; j <= H[i]; j++ {
		       fmt.Printf("-")
		   }
		   fmt.Printf("%d|\n",H[i])
		*/
		for H[i] < rectangles_bellow.Peek() {
			rectangles_bellow.Pop()
			if H[i] > rectangles_bellow.Peek() {
				rectangles_bellow.Push(H[i])
				rectangles += 1
			}
		}

		if H[i] > rectangles_bellow.Peek() {
			rectangles += 1
			rectangles_bellow.Push(H[i])
		}
		/*
		   fmt.Println(rectangles_bellow)
		   fmt.Printf("using %d rectangles\n",rectangles)
		*/
	}

	return rectangles
}
