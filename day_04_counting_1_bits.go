package main

import (
	"fmt"
)

func count_of_1_bit(n int) int {

	count := 0
	for n != 0 {

		n &= (n - 1)
		count++

	}
	return count
}

func main() {

	fmt.Println("count", count_of_1_bit(1001))

}
