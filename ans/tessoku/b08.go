package main

import (
	"fmt"
)

func main() {
	max := 1500

	var N int
	fmt.Scan(&N)

	S := make([][]int, max+1)
	for i := 0; i <= max; i++ {
		S[i] = make([]int, max+1)
	}

	X := make([]int, N+1)
	Y := make([]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Scan(&X[i], &Y[i])
		S[X[i]][Y[i]]++
	}

	for i := 1; i <= max; i++ {
		for j := 1; j <= max; j++ {
			S[i][j] += S[i-1][j]
		}
	}

	for j := 1; j <= max; j++ {
		for i := 1; i <= max; i++ {
			S[i][j] += S[i][j-1]
		}
	}

	var Q int
	fmt.Scan(&Q)

	for i := 0; i < Q; i++ {
		var A, B, C, D int
		fmt.Scan(&A, &B, &C, &D)
		fmt.Println(S[C][D] + S[A-1][B-1] - S[A-1][D] - S[C][B-1])
	}
}
