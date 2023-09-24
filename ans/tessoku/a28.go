package main

import "fmt"

func main() {
	ans := 0

	var N int
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		var T string
		var A int
		fmt.Scan(&T, &A)

		switch T {
		case "+":
			ans += A
		case "-":
			ans -= A
		case "*":
			ans *= A
		}

		if ans < 0 {
			ans += 10000
		}

		ans %= 10000

		fmt.Println(ans)
	}
}
