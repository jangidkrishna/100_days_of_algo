package main

import "fmt"

func sort(arr1 []int) []int {
	if len(arr1) <= 1 {
		return arr1
	}

	l, r := split(arr1)
	l = sort(l)
	r = sort(r)
	return merge(l, r)
}

func split(arr1 []int) ([]int, []int) {
	return arr1[0 : len(arr1)/2], arr1[len(arr1)/2:]
}

func merge(arr1, arr2 []int) []int {
	arr := make([]int, len(arr1)+len(arr2))

	j, k := 0, 0

	for i := 0; i < len(arr); i++ {
		if j >= len(arr1) {
			arr[i] = arr2[k]
			k++
			continue
		} else if k >= len(arr2) {
			arr[i] = arr1[j]
			j++
			continue
		}
		if arr1[j] > arr2[k] {
			arr[i] = arr2[k]
			k++
		} else {
			arr[i] = arr1[j]
			j++
		}
	}

	return arr
}

func main() {
	arr1 := []int{54, 33, 45, 3, 23, 213, 5, 43, 645, 7, 8}
	fmt.Println(sort(arr1))
}
