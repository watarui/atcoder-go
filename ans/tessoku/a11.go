package main

import "fmt"

func binarySearch(A []int, X int) int {
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
	return -1
}

func main() {
	var N, X int
	fmt.Scan(&N, &X)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	fmt.Println(binarySearch(A, X) + 1)
}
