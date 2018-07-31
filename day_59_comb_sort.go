package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generate_slice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func comb_sort(arr []int) {
	var (
		n       = len(arr)
		gap     = len(arr)
		shrink  = 1.3
		swapped = true
	)

	for swapped {
		swapped = false
		gap = int(float64(gap) / shrink)
		if gap < 1 {
			gap = 1
		}
		for i := 0; i+gap < n; i++ {
			if arr[i] > arr[i+gap] {
				arr[i+gap], arr[i] = arr[i], arr[i+gap]
				swapped = true
			}
		}
	}
}

func main() {

	slice := generate_slice(20)
	comb_sort(slice)
	fmt.Println(slice)
}
