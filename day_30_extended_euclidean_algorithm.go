package main

import "fmt"

func extended_euclidean_algorithm(a int, b int, x int, y int) int {

	if a == 0 {
		x = 0
		y = 1
		return b
	}

	temp_x1 := 1
	temp_y1 := 1

	gcd := extended_euclidean_algorithm(b%a, a, temp_x1, temp_y1)

	x = temp_y1 - (b/a)*temp_x1
	y = temp_x1

	return gcd
}

func main() {

	fmt.Println(extended_euclidean_algorithm(35, 15, 1, 1))

}
