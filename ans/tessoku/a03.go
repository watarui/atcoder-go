package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N, K int
	fmt.Scanf("%d %d", &N, &K)

	reader1 := bufio.NewReader(os.Stdin)
	input1, _ := reader1.ReadString('\n')
	input1 = strings.TrimSpace(input1)
	Ps := strings.Split(input1, " ")

	reader2 := bufio.NewReader(os.Stdin)
	input2, _ := reader2.ReadString('\n')
	input2 = strings.TrimSpace(input2)
	Qs := strings.Split(input2, " ")

	for _, P := range Ps {
		p, _ := strconv.Atoi(P)
		for _, Q := range Qs {
			q, _ := strconv.Atoi(Q)
			if p+q == K {
				fmt.Print("Yes")
				return
			}
		}
	}

	fmt.Print("No")
}
