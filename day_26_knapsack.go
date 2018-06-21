package main

import "fmt"

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func knapsack(value, weight []int, max_w int) int {
	d_p := make([][]int, len(value)+1)

	for i := range d_p {
		d_p[i] = make([]int, max_w+1)
		for j := range d_p[i] {
			d_p[i][j] = 0
		}
	}

	for i := 1; i <= len(value); i++ {
		for j := 1; j <= max_w; j++ {
			if weight[i-1] <= j {
				d_p[i][j] = max(d_p[i-1][j], value[i-1]+d_p[i-1][j-weight[i-1]])
			} else {
				d_p[i][j] = d_p[i-1][j]
			}

		}
	}

	return d_p[len(value)][max_w]
}

func main() {
	val := []int{58, 56, 656, 64, 54, 84, 1, 8, 48}
	weight := []int{12,32,45,5,65,65,43,32,4}
	max := 90
	fmt.Println( knapsack(val, weight,max))
}
