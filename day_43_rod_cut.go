package main

import (
	"fmt"
	"math"
)

const min float64 = -1 << 63

func rod_cut(price []float64, n int) float64 {
	val := []float64{0}
	val = append(val, price...)

	for i := 1; i < n+1; i++ {
		var max = min
		for j := 0; j < i; j++ {
			max = math.Max(max, price[j]+val[i-j-1])
		}
		val[i] = max
	}
	return val[n]
}

func main() {
	arr := []float64{1, 2, 3, 5, 8, 10, 11, 12}
	var size = len(arr)
	fmt.Println(rod_cut(arr, size))
}
