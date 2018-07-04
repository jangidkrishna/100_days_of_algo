package main

import "fmt"

func swap(x int, y int) (int, int) {

	s := (x - y) >> 31
	return y*(1+s) - x*s, x*(1+s) - y*s

}

func main() {

	x, y := swap(15, 13)
	fmt.Printf("swap 15 and 13 : %v & %v", x, y)

}
