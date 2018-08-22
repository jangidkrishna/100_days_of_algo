package main

import (
	"fmt"
)

func search(arr []int, x int) int {
	for i, n := range arr {
		if n == x {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{67, 342, 432, 423, 4, 23, 55, 6, 5457, 56, 867, 8, 6}
	find := 67
	index := search(arr, find)

	fmt.Println(index)

}
