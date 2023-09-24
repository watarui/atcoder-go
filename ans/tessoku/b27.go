package main

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	if a == 0 {
		return b
	}
	if a >= b {
		return gcd(a%b, b)
	} else {
		return gcd(a, b%a)
	}
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	fmt.Println(lcm(A, B))
}
