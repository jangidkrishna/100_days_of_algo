package main

import "fmt"

func coin_change(arr []int, limit int) int {

	x := make([]int, limit+1)
	x[0] = 1

	for _, v := range arr {
		for i := 1; i <= limit; i++ {
			if i >= v {
				x[i] += x[i-v]
			}
		}
	}

	return x[limit]
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	limit := 5

	way := coin_change(arr, limit)
	fmt.Printf("%d ways to combine %d from %v\n", way, limit, arr)
}
