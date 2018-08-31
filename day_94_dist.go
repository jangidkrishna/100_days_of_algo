package main

import (
	"fmt"
	"math"
)

type vector struct {
	x, y float64
}

func dist(a, b vector) float64 {
	return math.Sqrt((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y))
}

func main() {
	a := vector{x: 0, y: 0}
	b := vector{x: 5, y: 5}
	fmt.Println(dist(a, b))
}
