package main

import (
	"fmt"
	"math"
)

func isPrime(n, div int) bool {
	if div == 1 {
		return true
	}
	if n%div == 0 {
		return false
	} else {
		return isPrime(n, div-1)
	}
}

func isPrime2(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var Q int
	fmt.Scan(&Q)

	for i := 0; i < Q; i++ {
		var X int
		fmt.Scan(&X)

		if isPrime(X, int(math.Sqrt(float64(X)))) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
