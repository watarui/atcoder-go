package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	var N int
	fmt.Scan(&N)

	dp := make([]int, N+1)

	A := make([]int, N+1)
	for i := 2; i <= N; i++ {
		fmt.Scan(&A[i])
	}

	B := make([]int, N+1)
	for i := 3; i <= N; i++ {
		fmt.Scan(&B[i])
	}

	dp[1] = 0
	dp[2] = A[2]
	for i := 3; i <= N; i++ {
		dp[i] = min(dp[i-1]+A[i], dp[i-2]+B[i])
	}

	fmt.Println(dp[N])
}
