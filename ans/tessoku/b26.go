package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	deleted := make([]bool, N+1)
	for i := 2; i <= N; i++ {
		deleted[i] = false
	}

	for i := 2; i*i <= N; i++ {
		if deleted[i] {
			continue
		}
		for j := i * 2; j <= N; j += i {
			deleted[j] = true
		}
	}

	for i := 2; i <= N; i++ {
		if !deleted[i] {
			fmt.Println(i)
		}
	}
}
