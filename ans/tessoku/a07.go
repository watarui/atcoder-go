package main

import (
	"fmt"
)

func main() {
	var D int
	fmt.Scan(&D)
	var N int
	fmt.Scan(&N)

	A := make([]int, D+1)
	for i := 0; i < D+1; i++ {
		A[i] = 0
	}

	B := make([]int, D+1)
	for i := 0; i < D+1; i++ {
		B[i] = 0
	}

	for i := 0; i < N; i++ {
		var L, R int
		fmt.Scan(&L, &R)

		B[L]++
		if R+1 <= D {
			B[R+1]--
		}
	}

	for i := 1; i <= D; i++ {
		A[i] = A[i-1] + B[i]
	}

	for i := 1; i <= D; i++ {
		fmt.Println(A[i])
	}
}
