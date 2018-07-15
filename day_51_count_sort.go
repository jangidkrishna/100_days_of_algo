package main

import "fmt"

const max_uint = ^uint(0)
const max_int = int(max_uint >> 1)
const min_int = -max_int - 1

func countingSort(arr []int) {
	max_num := min_int
	min_num := max_int
	for _, v := range arr {
		if v > max_num {
			max_num = v
		}
		if v < min_num {
			min_num = v
		}
	}

	count := make([]int, max_num-min_num+1)

	for _, x := range arr {
		count[x-min_num]++
	}
	index := 0
	for i, c := range count {
		for ; c > 0; c-- {
			arr[index] = i + min_num
			index++
		}
	}
}

func main() {
	var arr = []int{987, 84, 894, 651, 984, 651, 68, 4351651, 351, 684, 35131, 35, 1}
	countingSort(arr)
	fmt.Println(arr)
}
