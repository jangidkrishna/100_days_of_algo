package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func binary_search(arr []int, l, r, x int) int {
	if r >= l {
		mid := int(l + (r-l)/2)

		if arr[mid] == x {
			return mid
		}

		if arr[mid] > x {
			return (binary_search(arr, l, mid-1, x))
		}

		return binary_search(arr, mid+1, r, x)
	}
	return -1
}

func exp_search(arr []int, n, x int) int {
	if arr[0] == x {
		return 0
	}
	i := 1
	for i < n && arr[i] <= x {
		i = i * 2
	}

	return binary_search(arr, i/2, min(i, n), x)

}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 546654, 8989989898}
	fmt.Println(exp_search(arr, len(arr)-1, 9))
}
