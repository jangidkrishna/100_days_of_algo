package main

import "fmt"

func pangram(str string) bool {
	var missing uint32 = (1 << 26) - 1
	for _, c := range str {
		var index uint32
		if 'a' <= c && c <= 'z' {
			index = uint32(c - 'a')
		} else if 'A' <= c && c <= 'Z' {
			index = uint32(c - 'A')
		} else {
			continue
		}

		missing &^= 1 << index
		if missing == 0 {
			return true
		}
	}
	return false
}

func main() {
	for _, str := range []string{
		"News is information about current events.",
		"Not a pangram.",
	} {
		if pangram(str) {
			fmt.Println("Yes:", str)
		} else {
			fmt.Println("No: ", str)
		}
	}
}
