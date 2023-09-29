package main

import (
	"fmt"
	"math"
)

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func reverse(s []int) []int {
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
	return s
}

func restore(current int, dp, h, V []int, f func(int) int) []int {
	V = append(V, current)
	if current == 1 {
		return V
	}
	if dp[current-1]+f(h[current]-h[current-1]) == dp[current] {
		return restore(current-1, dp, h, V, f)
	} else {
		return restore(current-2, dp, h, V, f)
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

	// 復元
	ans := reverse(restore(N, dp, h, []int{}, abs))

	al := len(ans)
	fmt.Println(al)
	for i := 0; i < al; i++ {
		if i >= 1 {
			fmt.Print(" ")
		}
		fmt.Print(ans[i])
	}
}
