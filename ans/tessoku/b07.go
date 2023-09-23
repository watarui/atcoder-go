package main

import (
	"fmt"
)

func main() {
	var T int // 閉店時刻
	fmt.Scan(&T)
	var N int // 従業員の人数
	fmt.Scan(&N)

	A := make([]int, T+1)
	for i := 0; i < T+1; i++ {
		A[i] = 0
	}

	B := make([]int, T+1)
	for i := 0; i < T+1; i++ {
		B[i] = 0
	}

	for i := 0; i < N; i++ {
		var L, R int // 出勤時刻、退勤時刻
		fmt.Scan(&L, &R)

		B[L]++
		B[R]--
	}

	for i := 1; i <= T; i++ {
		A[i] = A[i-1] + B[i-1]
	}

	for i := 1; i <= T; i++ {
		fmt.Println(A[i])
	}
}
