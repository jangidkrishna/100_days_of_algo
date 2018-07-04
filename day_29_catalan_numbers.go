package main

import (
	"fmt"
)

func bionomial_coeff(n uint64, k uint64) uint64 {
	res := uint64(1)

	if k > (n - k) {
		k = n - k
	}

	for i := uint64(0); i < k; i++ {
		res *= uint64(n - i)
		res /= uint64(1 + i)
	}

	return res
}

func catalan(n uint64) uint64 {
	c := bionomial_coeff(2*n, n)

	return (c / (n + 1))
}

func main() {
	fmt.Println(catalan(10))
}
