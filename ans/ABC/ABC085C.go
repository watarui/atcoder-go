package main

import (
	"fmt"
)

func main() {
	var N, Y int
	fmt.Scanf("%d %d", &N, &Y)

	for x := 0; x <= N; x++ {
		for y := 0; y <= N-x; y++ {
			z := N - x - y
			if 10000*x+5000*y+1000*z == Y {
				fmt.Printf("%d %d %d", x, y, z)
				return
			}
		}
	}
	fmt.Printf("-1 -1 -1")
}
