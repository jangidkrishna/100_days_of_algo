package main

import "fmt"

const mod = 1000000007

func fast_power(base int64, power int64) int64 {
	var result int64
	result = 1
	for power > 0 {

		if (power % 2) == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		power = power / 2
	}
	return result
}

func main() {
	fmt.Println(fast_power(2, 100))
}
