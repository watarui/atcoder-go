package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 各桁の和
func sum(i int) int {
	sum := 0
	for i > 0 {
		sum += i % 10
		i /= 10
	}
	return sum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	inputs := strings.Split(input, " ")

	n, _ := strconv.Atoi(inputs[0])
	a, _ := strconv.Atoi(inputs[1])
	b, _ := strconv.Atoi(inputs[2])

	total := 0
	for i := 1; i <= n; i++ {
		sum := sum(i)
		if sum >= a && sum <= b {
			total += i
		}
	}

	fmt.Printf("%d", total)
}
