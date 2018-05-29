package main

import (
	"fmt"
	"sort"
)

func permute(arr []int) {

	n := len(arr)
	var i, j int

	for i = n - 1; i > 0; i-- {
		if arr[i] > arr[i-1] {
			break
		}
	}

	x := arr[i-1]
	smallest := i

	for j = i; j < n; j++ {
		if arr[j] > x && arr[j] < arr[smallest] {
			smallest = j
		}
	}

	arr[smallest], arr[i-1] = arr[i-1], arr[smallest]

	sort.IntSlice.Sort(arr[i:])

	fmt.Println(arr)

}

func main() {
	arr1 := []int{5, 3, 4, 9, 7, 6}
	permute(arr1)

}
