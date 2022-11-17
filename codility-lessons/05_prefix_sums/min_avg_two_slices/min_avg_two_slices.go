package solution

// you can also use imports, for example:
import "math"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	size := len(A)
	avgs2 := make([]float64, size)
	avgs3 := make([]float64, size)
	minavgPos := 0
	curravg, minavg := math.MaxFloat32, math.MaxFloat32

	if size <= 2 {
		return 0
	}

	for j := 0; j < size-1; j++ {
		if j <= size-3 {
			avgs3[j] = float64(A[j]+A[j+1]+A[j+2]) / 3
		} else {
			avgs3[j] = math.MaxFloat32
		}
		if j <= size-2 {
			avgs2[j] = float64(A[j]+A[j+1]) / 2
		} else {
			avgs2[j] = math.MaxFloat32
		}
	}

	for k := 0; k < size-1; k++ {
		if avgs2[k] < avgs3[k] {
			curravg = avgs2[k]
		} else {
			curravg = avgs3[k]
		}
		if curravg < minavg {
			minavg = curravg
			minavgPos = k
		}
	}

	return minavgPos
}
