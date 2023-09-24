package main

import (
	"fmt"
	"sort"
)

func binarySearchLeft(A []int, X int) int {
	left := 0
	right := len(A) - 1
	for left <= right {
		mid := (left + right) / 2
		if A[mid] == X {
			return mid
		} else if A[mid] < X {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// // [min, ... , max]
// func sort(A []int) []int {
// 	for i := 0; i < len(A); i++ {
// 		for j := i + 1; j < len(A); j++ {
// 			if A[i] > A[j] {
// 				A[i], A[j] = A[j], A[i]
// 			}
// 		}
// 	}
// 	return A
// }

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
	sort.Slice(A, func(i, j int) bool { return A[i] < A[j] })

	var Q int
	fmt.Scan(&Q)

	for i := 0; i < Q; i++ {
		var X int
		fmt.Scan(&X)
		fmt.Println(binarySearchLeft(A, X))
	}
}
