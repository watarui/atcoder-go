package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func gcd(a, b int) int {
	x := 0
	for i := 1; i <= min(a, b); i++ {
		if a%i == 0 && b%i == 0 {
			x = i
		}
	}
	return x
}

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	fmt.Println(gcd(A, B))
}
