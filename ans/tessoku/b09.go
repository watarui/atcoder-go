package main

import (
	"fmt"
)

func main() {
	max := 1509

	var N int
	fmt.Scan(&N)

	P := make([][]int, max+1)
	for i := 0; i <= max; i++ {
		P[i] = make([]int, max+1)
	}

	for i := 0; i <= max; i++ {
		for j := 0; j <= max; j++ {
			P[i][j] = 0
		}
	}

	for i := 1; i <= N; i++ {
		var A, B, C, D int
		fmt.Scan(&A, &B, &C, &D)

		P[A][B]++
		P[A][D]--
		P[C][B]--
		P[C][D]++
	}

	for i := 0; i <= max; i++ {
		for j := 1; j <= max; j++ {
			P[i][j] += P[i][j-1]
		}
	}

	for j := 0; j <= max; j++ {
		for i := 1; i <= max; i++ {
			P[i][j] += P[i-1][j]
		}
	}

	ans := 0
	for j := 0; j <= max; j++ {
		for i := 0; i <= max; i++ {
			if P[i][j] > 0 {
				ans++
			}
		}
	}

	fmt.Println(ans)
}
