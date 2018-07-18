package main

import (
	"fmt"
	"math/rand"
)

func quick_select(arr []int, index int) int {
	return selec(arr, 0, len(arr)-1, index)
}

func partition(arr []int, l int, r int, pivot int) int {
	val := arr[pivot]
	arr[pivot], arr[r] = arr[r], arr[pivot]
	store := l

	for i := l; i < r; i++ {
		if arr[i] < val {
			arr[store], arr[i] = arr[i], arr[store]
			store += 1
		}
	}
	arr[r], arr[store] = arr[store], arr[r]

	return store
}

func selec(arr []int, l int, r int, k int) int {

	if l == r {
		return arr[l]
	}

	pivot := rand.Intn(r)
	pivot = partition(arr, l, r, pivot)

	if k == pivot {
		return arr[k]
	} else if k < pivot {
		return selec(arr, l, pivot-1, k)
	} else {
		return selec(arr, pivot+1, r, k)
	}
}

func main() {
	arr := []int{1, 2, 3, 5, 6, 7, 8}
	fmt.Println(quick_select(arr, 4))
}
