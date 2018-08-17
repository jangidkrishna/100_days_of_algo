package main

import "fmt"

func combination(n int, k int) int {
	prod := 1

	if n < k {
		return 0
	}
	if k > n/2 {
		k = n - k
	}
	for i := 1; i <= k; i++ {
		prod = prod * (n - i + 1) / i
	}

	return prod
}

func pascal_triangle(rows int) [][]int {
	matrix := make([][]int, rows+1, rows+1)

	for y := 0; y <= rows; y++ {
		matrix[y] = make([]int, y+1, y+1)
	}

	for y := 0; y <= rows; y++ {
		matrix[y][0] = 1
		matrix[y][y] = 1
	}

	for y := 2; y <= rows; y++ {
		for x := 1; x <= y-1; x++ {
			matrix[y][x] = matrix[y-1][x-1] + matrix[y-1][x]
		}
	}

	return matrix
}

func main() {

	mat := pascal_triangle(10)

	for i := 0; i <= 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(mat[i][j], " ")
		}
		fmt.Print("\n")
	}
}
