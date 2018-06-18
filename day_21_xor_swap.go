package main

import "fmt"

func xor_swap(a *int, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}

func main() {
	a, b := 10, 15
	fmt.Println("before ", a, b)
	xor_swap(&a, &b)
	fmt.Println("after", a, b)
}
