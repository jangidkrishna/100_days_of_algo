package main

import (
	"fmt"
)

func insertionsort(arr []float64) {

	for j := 1; j < len(arr); j++ {
		key := arr[j]
		i := j - 1

		for i >= 0 && arr[i] > key {
			arr[i+1] = arr[i]
			i = i - 1
		}

		arr[i+1] = key
	}
}

func main() {

	arr1 := []float64{3, 8, 43256, 8998, 2341, 4, 2, 2, 546898, 42389, 34889}
	insertionsort(arr1)
	fmt.Println("Sorted array : ", arr1)

}
