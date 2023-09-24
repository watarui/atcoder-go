package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	a1 := (N / 3)
	a2 := (N / 5)
	a3 := (N / 15)
	fmt.Println(a1 + a2 - a3)
}
