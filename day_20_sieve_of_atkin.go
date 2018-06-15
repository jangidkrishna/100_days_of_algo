package main

import (
	"fmt"
)

func sieve_of_atkin(limit int) {
	if limit > 2 {
		fmt.Println("2 ")
	}
	if limit > 3 {
		fmt.Println("3 ")
	}

	prime := make([]bool, limit)

	for i := 0; i < limit; i++ {
		prime[i] = false
	}

	for x := 1; x*x < limit; x++ {
		for y := 1; y*y < limit; y++ {
			n := (4 * x * x) + (y * y)
			if n <= limit && (n%12 == 1 || n%12 == 5) {
				prime[n] = (prime[n] || true) && !(prime[n] && true)
			}
			n = (3 * x * x) + (y * y)
			if n <= limit && n%12 == 7 {
				prime[n] = (prime[n] || true) && !(prime[n] && true)
			}
			n = (3 * x * x) - (y * y)
			if x > y && n <= limit && n%12 == 11 {
				prime[n] = (prime[n] || true) && !(prime[n] && true)
			}
		}
	}

	for r := 5; r*r < limit; r++ {
		if prime[r] {
			for i := r * r; i < limit; i += r * r {
				prime[i] = false
			}
		}
	}

	for a := 5; a < limit; a++ {
		if prime[a] {
			fmt.Println(a, " ")
		}
	}

}

func main() {
	sieve_of_atkin(5000)
}
