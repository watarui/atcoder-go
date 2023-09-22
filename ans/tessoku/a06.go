package main

import (
	"fmt"
)

func main() {
	var N, Q int
	fmt.Scan(&N, &Q)

	a := make([]int, N+1)
	s := make([]int, N+1)

	for i := 1; i <= N; i++ {
		fmt.Scan(&a[i])
		s[i] = a[i] + s[i-1]
	}

	for i := 0; i < Q; i++ {
		var L, R int
		fmt.Scan(&L, &R)
		fmt.Println(s[R] - s[L-1])
	}
}
