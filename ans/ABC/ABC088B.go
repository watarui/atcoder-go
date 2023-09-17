package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// [max, ... ,min]
func sort(ints []int) []int {
	for i := 0; i < len(ints); i++ {
		for j := i + 1; j < len(ints); j++ {
			if ints[i] < ints[j] {
				ints[i], ints[j] = ints[j], ints[i]
			} else {
				continue
			}
		}
	}
	return ints
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	as := strings.Split(input, " ")

	ints := []int{}
	for _, i := range as {
		a, _ := strconv.Atoi(i)
		ints = append(ints, a)
	}

	ints = sort(ints)

	alice := 0
	bob := 0
	for k, v := range ints {
		if k%2 == 0 {
			alice += v
		} else {
			bob += v
		}
	}

	fmt.Printf("%d", alice-bob)
}
