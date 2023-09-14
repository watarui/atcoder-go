package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	inputs := strings.Split(input, " ")

	a := []int{}
	for _, i := range inputs {
		m, _ := strconv.Atoi(i)
		a = append(a, m)
	}

	cnt := 0
Loop:
	for {
		for k, v := range a {
			mod := v % 2
			if mod == 1 {
				break Loop
			} else {
				a[k] = v / 2
			}
			if n == k+1 {
				cnt++
			}
		}
	}

	fmt.Printf("%d", cnt)
}
