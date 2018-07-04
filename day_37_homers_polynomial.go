package main

import "fmt"

func evaluate(arr []float64, x float64) float64 {
	ret := float64(0)
	for i := range arr {
		ret = ret*x + arr[i]
	}
	return ret
}

func main() {
	arr := []float64{1, 342, 54, 35, 436, 56}
	x := []float64{1, 321, 4, 3254, 3}
	j := 4
	for j >= 0 {
		fmt.Println(evaluate(arr, x[j]))
		j--
	}
}
