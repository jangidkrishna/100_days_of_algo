package main

import (
	"fmt"
)

func mccarthy91(n int) int {

	if n > 100 {
		return n - 10
	}

	return mccarthy91(mccarthy91(n + 11))
}

func main() {
	n := 45
	fmt.Println("Value returned by mccarthy func if n = 45 :", mccarthy91(n))
	k := 156
	fmt.Println("Value returned by mccarthy func if n = 156 :", mccarthy91(k))
}
