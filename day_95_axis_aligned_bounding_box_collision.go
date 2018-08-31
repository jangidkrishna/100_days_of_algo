package main

import (
	"fmt"
	"math"
)

type vector struct {
	x, y float64
}

type shape struct {
	center        vector
	height, width float64
}

func check(a, b shape) bool {
	return (math.Abs(a.center.x-b.center.x)*2 < (a.width + b.width)) &&
		(math.Abs(a.center.y-b.center.y)*2 < (a.height + b.height))
}

func main() {
	A := shape{center: vector{0, 0}, height: 6, width: 9}
	B := shape{center: vector{3, 9}, height: 6, width: 4}

	fmt.Println(check(A, B))
}
