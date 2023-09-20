package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N, X int
	fmt.Scanf("%d %d", &N, &X)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	strings := strings.Split(input, " ")

	for _, s := range strings {
		i, _ := strconv.Atoi(s)
		if i == X {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
