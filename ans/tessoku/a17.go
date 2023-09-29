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

func reverse(s []int) []int {
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
	return s
}

func restore(current int, dp, A, B, V []int) []int {
	V = append(V, current)
	if current == 1 {
		return V
	}
	if dp[current-1]+A[current] == dp[current] {
		return restore(current-1, dp, A, B, V)
	} else {
		return restore(current-2, dp, A, B, V)
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

	// 復元
	ans := reverse(restore(N, dp, A, B, []int{}))

	al := len(ans)
	fmt.Println(al)
	for i := 0; i < al; i++ {
		if i >= 1 {
			fmt.Print(" ")
		}
		fmt.Print(ans[i])
	}
}
