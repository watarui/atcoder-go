package main

import "fmt"

func main() {
	var N, K int
	fmt.Scanf("%d %d", &N, &K)

	patterns := 0

	for red := 1; red <= N; red++ {
		for blue := 1; blue <= N; blue++ {
			white := K - red - blue
			if white >= 1 && white <= N {
				patterns++
			}
		}
	}

	fmt.Printf("%d", patterns)
}
