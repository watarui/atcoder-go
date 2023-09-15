package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	inputs := strings.Split(input, " ")

	n, _ := strconv.Atoi(inputs[0])
	a, _ := strconv.Atoi(inputs[1])
	b, _ := strconv.Atoi(inputs[2])

	sum := 0
	for i := 1; i <= n; i++ {
		// 1桁
		if i < 10 {
			if i >= a && i <= b {
				sum = sum + i
			}
		}
		// 2桁
		if i >= 10 && i < 100 {
			n1 := i % 10
			n2 := i / 10
			n_sum := n1 + n2
			if n_sum >= a && n_sum <= b {
				sum = sum + i
			}
		}
		// 3桁
		if i >= 100 && i < 1000 {
			n1 := i % 10
			n2 := (i / 10) % 10
			n3 := i / 100
			n_sum := n1 + n2 + n3
			if n_sum >= a && n_sum <= b {
				sum = sum + i
			}
		}
		// 4桁
		if i >= 1000 && i < 10000 {
			n1 := i % 10
			n2 := (i / 10) % 10
			n3 := (i / 100) % 10
			n4 := i / 1000
			n_sum := n1 + n2 + n3 + n4
			if n_sum >= a && n_sum <= b {
				sum = sum + i
			}
		}
		// 5桁
		if i == 10000 {
			n_sum := 1
			if n_sum >= a && n_sum <= b {
				sum = sum + i
			}
		}
	}

	fmt.Printf("%d", sum)
}
