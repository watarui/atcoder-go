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

func gcd2(a, b int) int {
	if b == 0 {
		return a
	}
	if a == 0 {
		return b
	}
	if a >= b {
		return gcd2(a%b, b)
	} else {
		return gcd2(a, b%a)
	}
}

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	fmt.Println(gcd2(A, B))
}
