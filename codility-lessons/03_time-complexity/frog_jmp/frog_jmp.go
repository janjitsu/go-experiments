package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(X int, Y int, D int) int {
	// write your code in Go 1.4
	distance := X
	jumps := 0

	if distance == Y {
		return 0
	}

	jumps = (Y - X) / D
	distance = X + D*jumps

	if distance < Y {
		distance += D
		jumps += 1
	}

	return jumps
}
