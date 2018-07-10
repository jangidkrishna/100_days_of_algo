package main

import "fmt"

func shake_sort(arr []int) {
	l := 0
	r := len(arr) - 1
	for l <= r {
		new_l := r
		new_r := l
		swap := false
		for i := l; i < r; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swap = true
				new_r = i
			}
		}
		if !swap {
			break
		}
		r = new_r
		swap = false
		for i := r; i >= l; i-- {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swap = true
				new_l = i
			}
		}
		if !swap {
			break
		}
		l = new_l
	}
}

func main() {
	arr := [10]int{234, 23, 65, 778, 679, 879, 789, 1, 54312, 312}
	shake_sort(arr[:])
	fmt.Println(arr)
}
