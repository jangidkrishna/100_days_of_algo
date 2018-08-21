package main

import (
	"fmt"
	"math"
)

func jump(arr []int, x int) int {
	var l, r int
	length := len(arr)
	jump := int(math.Sqrt(float64(length)))

	for l < length && arr[l] <= x {
		r = min(length-1, l+jump)
		if arr[l] <= x && arr[r] >= x {
			break
		}
		l += jump
	}

	if l >= length || arr[l] > x {
		return -1
	}

	r = min(length-1, r)
	index := l

	for index <= r && arr[index] <= x {
		if arr[index] == x {
			return index
		}
		index++
	}

	return -1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	arr := []int{0, 2, 5, 8, 11, 24, 56, 69, 98}
	fmt.Println(jump(arr, 8))
}
