package main

import (
	"fmt"
	"strings"
)

func palindrome(str string) bool {
	str = strings.ToLower(str)
	le := len(str)
	l := 0
	r := le - 1

	for l <= r {
		if str[l] != str[r] {
			return false
		}
		l = l + 1
		r = r - 1
	}
	return true
}

func main() {
	if palindrome("anna") {
		fmt.Println("Anna is palindrome")
	}
}
