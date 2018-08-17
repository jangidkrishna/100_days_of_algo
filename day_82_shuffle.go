package main

import "fmt"
import "math/rand"

func shuffle(data []int) {
	n := len(data)
	for i := 0; i < n; i++ {
		k := rand.Intn(n)
		data[i], data[k] = data[k], data[i]
	}
	fmt.Println(data)
}

func main() {
	arr := []int{79, 10, 73, 1, 68, 56, 3}
	for i := 0; i < 10; i++ {
		shuffle(arr)
	}
}
