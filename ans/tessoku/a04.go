package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)

	fmt.Print(strings.Join([]string{
		strconv.Itoa((N / int(math.Pow(2, 9))) % 2),
		strconv.Itoa((N / int(math.Pow(2, 8))) % 2),
		strconv.Itoa((N / int(math.Pow(2, 7))) % 2),
		strconv.Itoa((N / int(math.Pow(2, 6))) % 2),
		strconv.Itoa((N / int(math.Pow(2, 5))) % 2),
		strconv.Itoa((N / int(math.Pow(2, 4))) % 2),
		strconv.Itoa((N / int(math.Pow(2, 3))) % 2),
		strconv.Itoa((N / int(math.Pow(2, 2))) % 2),
		strconv.Itoa((N / int(math.Pow(2, 1))) % 2),
		strconv.Itoa((N / int(math.Pow(2, 0))) % 2),
	}, ""))
}
