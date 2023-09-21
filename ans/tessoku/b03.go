package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	A := strings.Split(input, " ")

	for i, xs := range A {
		x, _ := strconv.Atoi(xs)
		for _, ys := range A[i+1:] {
			y, _ := strconv.Atoi(ys)
			for _, zs := range A[i+2:] {
				z, _ := strconv.Atoi(zs)
				if x+y+z == 1000 {
					fmt.Print("Yes")
					return
				}
			}
		}
	}

	fmt.Print("No")
}
