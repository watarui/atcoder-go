package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N+1)
	wins := make([]int, N+1)
	losses := make([]int, N+1)

	wins[0] = 0
	losses[0] = 0
	for i := 1; i <= N; i++ {
		fmt.Scan(&A[i])
		if A[i] == 0 {
			wins[i] = wins[i-1]
			losses[i] = losses[i-1] + 1
		} else {
			wins[i] = wins[i-1] + 1
			losses[i] = losses[i-1]
		}
	}

	var Q int
	fmt.Scan(&Q)

	for i := 0; i < Q; i++ {
		var L, R int
		fmt.Scan(&L, &R)
		win := wins[R] - wins[L-1]
		lose := losses[R] - losses[L-1]
		if win > lose {
			fmt.Println("win")
		} else if win == lose {
			fmt.Println("draw")
		} else {
			fmt.Println("lose")
		}
	}
}
