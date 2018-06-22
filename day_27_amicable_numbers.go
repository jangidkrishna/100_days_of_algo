package main

import "fmt"

func prime_factors(n int) (pfs []int) {
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}
	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

func power(p, i int) int {
	result := 1
	for j := 0; j < i; j++ {
		result *= p
	}
	return result
}

func sum_of_proper_divisors(n int) int {
	pfs := prime_factors(n)
	m := make(map[int]int)
	for _, prime := range pfs {
		_, ok := m[prime]
		if ok {
			m[prime] += 1
		} else {
			m[prime] = 1
		}
	}

	sum_of_all_factors := 1
	for prime, exponents := range m {
		sum_of_all_factors *= (power(prime, exponents+1) - 1) / (prime - 1)
	}
	return sum_of_all_factors - n
}

func amicable_numbers_under_10000() (amicables []int) {
	for i := 3; i < 10000; i++ {
		s := sum_of_proper_divisors(i)
		if s == i {
			continue
		}
		if sum_of_proper_divisors(s) == i {
			amicables = append(amicables, i)
		}
	}
	return
}

func main() {
	amicables := amicable_numbers_under_10000()
	fmt.Println(amicables)

	sum := 0
	for i := 0; i < len(amicables); i++ {
		sum += amicables[i]
	}
	fmt.Println(sum)
}
