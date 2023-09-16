package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func uniq(ints []int) []int {
	m := make(map[int]bool)
	uniq := []int{}

	for _, v := range ints {
		if !m[v] {
			m[v] = true
			uniq = append(uniq, v)
		}
	}
	return uniq
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var ss []string
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		ss = append(ss, scanner.Text())
	}

	ds := []int{}
	for _, i := range ss {
		m, _ := strconv.Atoi(i)
		ds = append(ds, m)
	}

	fmt.Println(len(uniq(ds)))
}
