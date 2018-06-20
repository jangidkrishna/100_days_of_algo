package main

import "fmt"

func two_sum(arr []float64, sum float64) (float64, float64) {
	n := len(arr)
	m := make(map[float64]bool, 0)
	for i := 0; i < n; i++ {
		m[arr[i]] = true
	}
	for i := 0; i < n; i++ {
		if m[sum-arr[i]] {
			return arr[i], sum - arr[i]
		}
	}
	return 0.0, 0.0
}

func main() {
	arr := []float64{58, 56, 656, 64, 54, 84, 1, 8, 48}
	fmt.Println(two_sum(arr, 85))
}
