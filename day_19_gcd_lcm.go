package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		c := b
		b = a % b
		a = c
	}
	return a
}

func lcm(a, b int) int {
	res := a * b / gcd(a, b)
	return res
}

func main() {
	fmt.Println(gcd(12, 8))
	fmt.Println(lcm(32, 94))
}
