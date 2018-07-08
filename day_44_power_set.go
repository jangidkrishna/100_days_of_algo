package main

import "fmt"

func power_set(arr, rec []int, n int) {
	if n == len(arr) {
		fmt.Print("{")
		for i, v := range rec {
			if v == 1 {
				fmt.Print(" ", arr[i])
			}
		}
		fmt.Println(" }")
		return
	}

	rec[n] = 1
	power_set(arr, rec, n+1)
	rec[n] = 0
	power_set(arr, rec, n+1)

}

func main() {
	arr := []int{1, 2, 3}
	arr2 := make([]int, len(arr))

	power_set(arr, arr2, 0)
}
