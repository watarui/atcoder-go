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
	fmt.Scan(&n)

	pos := [][]int{{0, 0, 0}}
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()

		strings := strings.Split(scanner.Text(), " ")
		ints := []int{}
		for _, str := range strings {
			s, _ := strconv.Atoi(str)
			ints = append(ints, s)
		}
		pos = append(pos, ints)
	}

	for i, p := range pos {
		if i == 0 {
			continue
		}

		dt := p[0] - pos[i-1][0]
		dx := p[1] - pos[i-1][1]
		dy := p[2] - pos[i-1][2]
		if dx < 0 {
			dx = -dx
		}
		if dy < 0 {
			dy = -dy
		}

		if dt < dx+dy || (dt-dx-dy)%2 != 0 {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}
