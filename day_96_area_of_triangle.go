package main

import (
	"fmt"
	"math"
)

type vector struct {
	x, y float64
}

func area(a, b, c vector) float64 {
	return math.Abs((a.x-b.x)*(c.y-b.y)-(a.y-b.y)*(c.x-b.x)) / 2
}

func main() {
	a := vector{x: 0, y: 0}
	b := vector{x: 0, y: 5}
	c := vector{x: 5, y: 0}

	fmt.Println(area(a, b, c))

}
