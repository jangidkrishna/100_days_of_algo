package main

import (
	"fmt"
	"sort"
)

func median_of_medians(arr []int, k, r int) int {

	num := len(arr)
	if num < 10 {
		sort.Ints(arr)
		return arr[k-1]
	}
	med := (num + r - 1) / r

	medians := make([]int, med)

	for i := 0; i < med; i++ {
		v := (i * r) + r
		var arr []int
		if v >= num {
			arr = make([]int, len(arr[(i*r):]))
			copy(arr, arr[(i*r):])
		} else {
			arr = make([]int, r)
			copy(arr, arr[(i*r):v])
		}
		sort.Ints(arr)
		medians[i] = arr[len(arr)/2]
	}
	pivot := median_of_medians(medians, (len(medians)+1)/2, r)

	var leftSide, rightSide []int

	for i := range arr {
		if arr[i] < pivot {
			leftSide = append(leftSide, arr[i])
		} else if arr[i] > pivot {
			rightSide = append(rightSide, arr[i])
		}
	}

	switch {
	case k == (len(leftSide) + 1):
		return pivot
	case k <= len(leftSide):
		return median_of_medians(leftSide, k, r)
	default:
		return median_of_medians(rightSide, k-len(leftSide)-1, r)
	}
}

func main() {
	arr := []int{1, 3, 543, 55, 65, 76, 7, 67, 876, 9, 8, 756, 46, 345, 32, 4, 234, 32, 12, 54}
	sort.Ints(arr)

	for _, j := range []int{5, 10, 15, 20} {
		i := median_of_medians(arr, j, 5)
		fmt.Println(j, "smallest : ", i)
		v := arr[j-1]
		fmt.Println("arr[", j-1, "] = ", v)
	}
}
