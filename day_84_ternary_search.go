package main

import "fmt"

func ternary_search(arr []int, left, right, x int) int {
	if right >= left {
		vl := (right - left) / 3
		leftmid := left + vl
		rightmid := leftmid + vl

		if arr[leftmid] == x {
			return leftmid
		}

		if arr[rightmid] == x {
			return rightmid
		}

		if x < arr[leftmid] {
			return ternary_search(arr, left, leftmid, x)
		}

		if x < arr[rightmid] {
			return ternary_search(arr, rightmid, right, x)
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 7, 8, 9, 565, 8989, 656565}
	fmt.Println(ternary_search(arr, 0, len(arr)-1, 565))
}
