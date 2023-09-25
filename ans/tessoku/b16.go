package main

import (
	"fmt"
	"math"
)

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

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

	h := make([]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Scan(&h[i])
	}

	dp[1] = 0
	dp[2] = abs(h[2] - h[1])
	for i := 3; i <= N; i++ {
		dp[i] = min(dp[i-1]+abs(h[i]-h[i-1]), dp[i-2]+abs(h[i]-h[i-2]))
	}

	fmt.Println(dp[N])
}
