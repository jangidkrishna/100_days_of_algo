package main

import "fmt"

var sols int

func pos(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func place(lay []int, row, col int) bool {
	for i := 0; i < row; i++ {
		if lay[i] == col {
			return false
		}
		if pos(lay[i]-col) == pos(i-row) {
			return false
		}
	}

	return true
}

func n_queen(lay []int, row int) {
	if row == len(lay) {
		sols++
		fmt.Printf("\n%d :=\n\n", sols)
		for i := 0; i < len(lay); i++ {
			for j := 0; j < len(lay); j++ {
				if lay[i] == j {
					fmt.Printf("Q")
				} else {
					fmt.Printf("-")
				}
			}
			fmt.Printf("\n")
		}
		return
	}

	for i := 0; i < len(lay); i++ {
		if place(lay, row, i) {
			lay[row] = i
			n_queen(lay, row+1)
		}
	}

}

func main() {
	board := make([]int, 4)
	sols = 0
	n_queen(board, 0)
}
