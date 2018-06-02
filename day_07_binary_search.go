package main

import (
	"fmt"
)

func binary_search(arr []int, start int, end int, key int) int {

	if end >= start {

		mid := start + (end-start)/2

		if arr[mid] == key {
			return mid
		}
		if arr[mid] > key {
			return binary_search(arr, start, mid-1, key)
		}

		return binary_search(arr, mid+1, end, key)

	}

	return -1
}

func main() {

	arr := []int{2, 3, 4, 10, 40}
	result := binary_search(arr, 0, len(arr), 10)
	if result == -1 {
		fmt.Println("Element is not present in the array")
	} else {
		fmt.Println("Element is at pos", result)
	}

}
