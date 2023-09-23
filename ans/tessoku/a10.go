package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N+1)
	A[0] = 0
	for i := 1; i <= N; i++ {
		fmt.Scan(&A[i])
	}

	var D int
	fmt.Scan(&D)

	L := make([]int, D+1)
	R := make([]int, D+1)
	L[0] = 0
	R[0] = 0
	for i := 1; i <= D; i++ {
		fmt.Scan(&L[i], &R[i])
	}

	P := make([]int, N+1)
	Q := make([]int, N+1)
	P[0] = 0
	P[1] = A[1]
	for i := 2; i <= N; i++ {
		P[i] = max(P[i-1], A[i])
	}

	Q[0] = 0
	Q[N] = A[N]
	for i := N - 1; i > 0; i-- {
		Q[i] = max(Q[i+1], A[i])
	}

	for d := 1; d <= D; d++ {
		fmt.Println(max(P[L[d]-1], Q[R[d]+1]))
	}
}
