package main

import (
	"fmt"
)

func cocktail_sort(arr []int) {
	count := 1
	for count > 0 {
		count = 0
		for index := 0; index < len(arr)-1; index++ {
			if arr[index] > arr[index+1] {
				arr[index], arr[index+1] = arr[index+1], arr[index]
				count += 1
			}
		}
		for index := len(arr) - 1; index > 0; index-- {
			if arr[index] < arr[index-1] {
				arr[index], arr[index-1] = arr[index-1], arr[index]
				count += 1
			}
		}
	}
}

func main() {
	arr := []int{8,3,54,68,7,9,1,0}
	cocktail_sort(arr)
	fmt.Println(arr)
}
