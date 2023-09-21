package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)
	A := strings.Split(fmt.Sprintf("%08d", N), "")

	digit := 8
	ten := 0
	for i := 0; i < digit; i++ {
		a, _ := strconv.Atoi(A[i])
		ten += a * int(math.Pow(2, float64((digit-1)-i)))
	}

	fmt.Printf("%d", ten)
}
