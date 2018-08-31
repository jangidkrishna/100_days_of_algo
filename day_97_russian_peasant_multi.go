package main

import "fmt"

func multi(a, b int) int {
	res := 0

	for b > 0 {
		if (b & 1) != 0 {
			res += a
		}
		a = a << 1
		b = b >> 1
	}
	return res
}

func main() {
	fmt.Println(multi(24, 15))
}
