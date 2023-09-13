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
	params := strings.Split(input, "")

	s1, _ := strconv.Atoi(params[0])
	s2, _ := strconv.Atoi(params[1])
	s3, _ := strconv.Atoi(params[2])

	cnt := 0
	for _, s := range []int{s1, s2, s3} {
		if s == 1 {
			cnt++
		}
	}

	fmt.Printf("%d", cnt)
}
