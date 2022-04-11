package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(X int, A []int) int {
	// write your code in Go 1.4

	/*
		estou na posição inicial 1
		initial = 1
		quero chegar na posição final
		final = X
		e preciso que entre initial e final todos os números apareçam pelo menos uma vez, mas eu não tenho a ordem
		eu só preciso contar se esses números apareceram e a partir que posição todos eles eram maior que zero
	*/
	frog := 1
	counts := make([]int, X+1)
	jumps_needed := X - frog
	jumps_made := 0
	counts[A[0]] = 1
	earlyest := 0
	// count jumps
	for i := 1; i < len(A); i++ {
		counts[A[i]] += 1
		if A[i] <= X && counts[A[i]] == 1 {
			jumps_made += 1
		}
		//fmt.Printf("i: %d, A[i]: %d; jumps_needed: %d; jumps_made: %d\n",i, A[i], jumps_needed, jumps_made)
		//fmt.Printf("count: %v\n", counts)
		if jumps_made == jumps_needed {
			earlyest = i
			break
		}
	}

	//check if it's possible to jump
	for i := 1; i <= X; i++ {
		if counts[i] == 0 {
			return -1
		}
	}

	return earlyest
}
