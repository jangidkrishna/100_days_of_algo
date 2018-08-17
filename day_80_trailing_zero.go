package main

import "fmt"

func trailing_zeros(n int) int {
	count := 0
	for i := 5; n/i >= 1; i *= 5 {
		count += n / i
	}
	return count
}

func main() {
	fmt.Println(trailing_zeros(100))
}
