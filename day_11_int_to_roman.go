package main

import (
	"fmt"
)

func int_to_roman(n int) {

	mapping := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	var result string

	for k, v := range mapping {
		for n >= k {
			result = result + v
			n -= k
		}
	}
	fmt.Println(result)
}

func main() {
	int_to_roman(100)
}
