package main

import (
	"fmt"
)

func pos(n int, k int) int {
	if n <= 1 {
		return 0
	}
	return (pos(n-1, k) + k) % n
}

func main() {
	fmt.Println(pos(10, 2))
}
