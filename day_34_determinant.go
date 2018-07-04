package main

import (
	"fmt"
	"math"
)

func determinant(n int, mat [][]int) float64 {

	d := float64(0)
	var c, subi, i, j, subj int
	var submat [][]int

	if n == 2 {
		return (float64((mat[0][0] * mat[1][1]) - (mat[1][0] * mat[0][1])))
	} else {

		for c = 0; c < n; c++ {

			subi = 0
			for i = 1; i < n; i++ {

				subj = 0
				for j = 0; j < n; j++ {

					submat[subi][subj] = mat[i][j]
					subj++
				}
				subi++
			}
			d = d + float64((math.Pow(-1, float64(c)) * float64(mat[0][c]) * determinant(n-1, submat)))
		}
	}
	return d

}

func main() {
	a := [][]int{
		{0, 1},
		{4, 5},
	}
	fmt.Println(determinant(2, a))
}
