package main

import "fmt"

func stooge_sort(arr []int, l, r int) {
	if arr[l] > arr[r] {
		arr[l], arr[r] = arr[r], arr[l]
	}

	if (r - l) >= 2 {
		mid := (r - l + 1) / 3
		stooge_sort(arr, l, r-mid)
		stooge_sort(arr, l+mid, r)
		stooge_sort(arr, l, r-mid)

	}
}

func main() {
	arr := []int{234,1,5,6547,8678,9879}
	stooge_sort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
