package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func shuffle(arr []int) []int {
	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func bogo_sort(arr []int) []int {
	for {
		if sort.IntsAreSorted(arr) {
			return arr
		}
		arr = shuffle(arr[:])
	}
}

func main() {
	arr := []int{9, 89, 809, 809, 889}
	fmt.Println(bogo_sort(arr))
}
