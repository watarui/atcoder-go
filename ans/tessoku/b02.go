package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scanf("%d %d", &A, &B)

	divisor := []int{100, 50, 25, 20, 10, 5, 4, 2, 1}

	for i := A; i <= B; i++ {
		for _, v := range divisor {
			if i == v {
				fmt.Println("Yes")
				return
			}
		}
	}

	fmt.Println("No")
}
