package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	var ni int
	ni, _ = strconv.Atoi(os.Args[1])
	var count int = 0
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 10000; i++ {

		random_x := random.Float64()
		random_y := random.Float64()
		dist := random_x*random_x + random_y*random_y

		if dist <= 1 {
			count++
		}
	}

	pi := float64(count) / float64(ni) * float64(4)
	fmt.Printf("# of trails = %v , estimated pi %v", ni, pi)

}
