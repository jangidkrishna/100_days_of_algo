package main

import (
	"fmt"
)

func multiply(arr1 []int, arr2 []int, m int, n int) []int {

	prod := make([]int, m+n-1)

	for i := 0; i < m+n-1; i++ {
		prod[i] = 0
	}

	for i := 0; i < m; i++ {

		for j := 0; j < n; j++ {
			prod[i+j] += arr1[i] * arr2[j]
		}
	}
	return prod
}

func main() {

	arr1 := []int{5, 0, 10, 6}
	arr2 := []int{1, 2, 4}
	fmt.Println(multiply(arr1, arr2, len(arr1), len(arr2)))

}
