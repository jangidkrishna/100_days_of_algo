package main

import "fmt"

func rec(n, m int) int {
	if n == m {
		return n
	}
	if n == 0 || m == 0 {
		return n + m
	}

	if n%2 == 0 {
		if m%2 == 1 {
			return rec(n>>1, m)
		} else {
			return rec(n>>1, m>>1) << 1
		}
	}

	if m%2 == 0 {
		return rec(n, m>>1)
	}

	if n > m {
		return rec((n-m)>>1, m)
	}

	return rec((m-n)>>1, n)
}

func iterate(n, m int) int {
	if n == m {
		return n
	}
	if n == 0 || m == 0 {
		return n + m
	}

	shift := 0
	for shift = 0; ((n | m) & 1) == 0; shift++ {
		n >>= 1
		m >>= 1
	}

	for (n & 1) == 0 {
		n >>= 1
	}

	for m != 0 {
		for m&1 == 0 {
			m >>= 1
		}
		if n > m {
			m, n = n, m
		}
		m -= n
	}

	return n << uint32(shift)
}

func main() {

	fmt.Println(iterate(14, 7))
}
