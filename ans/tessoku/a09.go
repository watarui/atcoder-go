package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	max := 1509

	var H, W, N int
	fmt.Scan(&H, &W, &N)

	X := make([][]int, max+1)
	Z := make([][]int, max+1)
	for i := 0; i <= max; i++ {
		X[i] = make([]int, max+1)
		Z[i] = make([]int, max+1)
	}

	for i := 1; i <= N; i++ {
		var A, B, C, D int
		fmt.Scan(&A, &B, &C, &D)

		X[A][B]++
		X[A][D+1]--
		X[C+1][B]--
		X[C+1][D+1]++
	}

	for i := 0; i <= H; i++ {
		for j := 0; j <= W; j++ {
			Z[i][j] = 0
		}
	}

	for i := 1; i <= H; i++ {
		for j := 1; j <= W; j++ {
			Z[i][j] = Z[i][j-1] + X[i][j]
		}
	}

	for j := 1; j <= W; j++ {
		for i := 1; i <= H; i++ {
			Z[i][j] += Z[i-1][j]
		}
	}

	for i := 1; i <= H; i++ {
		xs := make([]string, W+1)
		for j := 1; j <= W; j++ {
			s := strconv.Itoa(Z[i][j])
			xs = append(xs, s)
		}
		fmt.Println(strings.Join(xs, " "))
	}
}
