package main

import (
	"fmt"
	"regexp"
)

func main() {
	var S string
	fmt.Scanf("%s", &S)

	if regexp.MustCompile(`^(dream|dreamer|erase|eraser)+$`).Match([]byte(S)) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
