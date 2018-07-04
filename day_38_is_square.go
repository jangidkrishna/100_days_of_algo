package main

import (
	"fmt"
	"math"
)

func is_square(x int) bool {

	if x < 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(x)))
	return sqrt*sqrt == x
}

func main() {
	fmt.Println(is_square(16))
}
