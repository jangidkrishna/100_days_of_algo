package main

import "fmt"

func newman_conway(number int) []int {
	arr := []int{}
	arr = append(arr, 0)
	arr = append(arr, 1)
	arr = append(arr, 1)

	for i := 3; i <= number; i++ {
		a := arr[arr[i-1]] + arr[i-arr[i-1]]
		arr = append(arr, a)
	}

	return arr
}

func main() {
	seq := newman_conway(10)
	fmt.Println(seq)
}
