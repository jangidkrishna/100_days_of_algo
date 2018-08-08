package main

import "fmt"

func tribonacci(n int) {
	a, b, c := 0, 1, 1
	res := 0
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	for i := 3; i < n; i++ {
		res = a + b + c
		a, b, c = b, c, res
		fmt.Println(res)
	}

}

func main() {
	tribonacci(5)
}
