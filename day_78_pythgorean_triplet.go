package main

import "fmt"

func pythagorean_tripplet(num int) {

	c := 0
	a := 0
	b := 0
	m := 2

	for c < num {
		for n := 1; n < m; n++ {
			a = m*m - n*n
			b = 2 * m * n
			c = m*m + n*n

			if c > num {
				break
			}

			fmt.Printf("%d %d %d\n", a, b, c)
		}
		m++
	}
}

func main() {
	pythagorean_tripplet(20)
}
