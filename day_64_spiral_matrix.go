package main

import (
	"fmt"
)

func spiral(n int) []int {
	left, top, right, bottom := 0, 0, n-1, n-1
	sz := n * n
	s := make([]int, sz)
	i := 0
	for left < right {
		for c := left; c <= right; c++ {
			s[top*n+c] = i
			i++
		}
		top++
		for r := top; r <= bottom; r++ {
			s[r*n+right] = i
			i++
		}
		right--
		if top == bottom {
			break
		}
		for c := right; c >= left; c-- {
			s[bottom*n+c] = i
			i++
		}
		bottom--
		for r := bottom; r >= top; r-- {
			s[r*n+left] = i
			i++
		}
		left++
	}
	s[top*n+left] = i

	return s
}

func main() {
	num := 5
	for i, draw := range spiral(num) {
		fmt.Printf("%*d ", 2, draw)
		if i%num == num-1 {
			fmt.Println("")
		}
	}
}
