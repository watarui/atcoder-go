package main

import (
	"fmt"
)

func main() {
	var H, W int // 行、列
	fmt.Scan(&H, &W)

	X := make([][]int, H+1)
	S := make([][]int, H+1)
	S[0] = make([]int, W+1)
	for i := 1; i <= H; i++ {
		X[i] = make([]int, W+1)
		S[i] = make([]int, W+1)
		for j := 1; j <= W; j++ {
			fmt.Scan(&X[i][j])
			S[i][j] = S[i][j-1] + X[i][j]
		}
	}

	for j := 1; j <= W; j++ {
		for i := 1; i <= H; i++ {
			S[i][j] += S[i-1][j]
		}
	}

	var Q int // 質問
	fmt.Scan(&Q)

	for i := 0; i < Q; i++ {
		var A, B, C, D int
		fmt.Scan(&A, &B, &C, &D)
		fmt.Println(S[C][D] + S[A-1][B-1] - S[A-1][D] - S[C][B-1])
	}
}
