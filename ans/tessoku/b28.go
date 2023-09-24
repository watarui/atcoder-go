package main

import "fmt"

func fibMod(n, mod int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return (fibMod(n-1, mod) + fibMod(n-2, mod)) % mod
	}
}

func main() {
	mod := 1000000007

	var N int
	fmt.Scan(&N)

	a := make([]int, N+1)
	a[1] = 1
	a[2] = 1
	for i := 3; i <= N; i++ {
		a[i] = (a[i-1] + a[i-2]) % mod
	}

	fmt.Println(a[N])
}
