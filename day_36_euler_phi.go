package main

import "fmt"

func euler_phi(n int) int {
	res := n
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			for n%i == 0 {
				n /= i
			}
			res *= (i - 1)
			res /= i
		}
	}
	if n != 1 {
		res *= (n - 1)
		res /= n
	}
	return res
}

func main() {
	fmt.Println(euler_phi(100))
}
