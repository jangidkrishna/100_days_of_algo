package main

import (
	"fmt"
	"math"
)

func mean(arr []float64) float64 {
	sum := 0.0
	n := len(arr)
	if n == 0 {
		return sum
	}
	for i := 0; i < n; i++ {
		sum = sum + arr[i]
	}

	return sum / float64(n)
}

func variance(arr []float64) float64 {
	variance := 0.0
	n := len(arr)
	if n == 0 {
		return variance
	}
	mean := mean(arr)
	for i := 0; i < n; i++ {
		variance += (arr[i] - mean) * (arr[i] - mean)
	}

	return variance / float64(n)
}

func std(arr []float64) float64 {
	return math.Sqrt(variance(arr))
}

func main() {
	arr := []float64{342, 42, 34, 324, 234, 32, 4234, 324}
	fmt.Println(std(arr))
}
