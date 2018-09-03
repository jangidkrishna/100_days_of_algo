package main

import "fmt"

func reverse(n int) {
	divisor, num := 2, n

	for 0 == num%divisor {
		num /= divisor
		divisor++
	}

	if num == 1 {
		fmt.Printf("%d is %d!\n", n, divisor-1)
		return
	}
	fmt.Printf(n)
	return
}

func main() {
	reverse(100)
}
