package main

import "fmt"

func delannoy_generator(n int, m int) int {
	d := 1
	if (n == 0) || (m == 0) {
		d = 1
	} else {
		d = delannoy_generator(n-1, m) + delannoy_generator(n, m-1) + delannoy_generator(n-1, m-1)
	}
	return d
}

func main() {
	fmt.Println(delannoy_generator(10, 20))
}
